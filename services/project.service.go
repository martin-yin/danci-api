package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"strings"
)

func GetProjectList(id uint) (projectList []model.Project, err error) {
	var teamList []model.Team
	err = global.GVA_DB.Preload("Admins", "id = ? ", 1).Preload("Projects").Model(&model.Team{}).Find(&teamList).Error
	for _, team := range teamList {
		for _, project := range team.Projects {
			projectList = append(projectList, project)
		}
	}
	return
}

func CreateProject(project model.Project) (projectInter model.Project, err error) {
	err = global.GVA_DB.Model(&model.Project{}).Create(&project).Error
	return project, err
}

func FindProject(projectName string) (isExist bool) {
	var project model.Project
	result := global.GVA_DB.Model(&model.Project{}).Where("project_name = ? ", projectName).First(&project)
	if result.RowsAffected != 0 {
		return true
	}
	return false
}

func GetProjectStatistics(startTime string, endTime string, monitorId string) (projectStatistics response.ProjectStatistics, err error) {
	err = global.GVA_DB.Model(&model.PageView{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(SqlWhereBuild("page_views"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
	return
}

func GetProjectHealthy(startTime string, endTime string, monitorIds string) (projectStatisticsList []response.ProjectStatistics, err error) {
	monitorIdealists := strings.Split(monitorIds, `,`)
	for _, monitorId := range monitorIdealists {
		var projectStatistics response.ProjectStatistics
		err = global.GVA_DB.Model(&model.PageView{}).Select("COUNT( DISTINCT user_id ) as uv, COUNT( DISTINCT id ) as pv").Where(SqlWhereBuild("page_views"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		// projectStatistics.Pv  当前这个项目总的pv
		err = global.GVA_DB.Model(&model.PageJsError{}).Select("COUNT( DISTINCT id ) as js_error").Where(SqlWhereBuild("page_js_errors"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		err = global.GVA_DB.Model(&model.PageResourceError{}).Select("COUNT( DISTINCT id ) as resources_error").Where(SqlWhereBuild("page_resource_errors"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		err = global.GVA_DB.Model(&model.PageHttp{}).Select("COUNT( DISTINCT id ) as http_error").Where(SqlWhereBuild("page_https"), startTime, endTime, monitorId).Scan(&projectStatistics).Error
		projectStatistics.JsError = Decimal(projectStatistics.JsError / projectStatistics.Pv)
		projectStatistics.ResourcesError = Decimal(projectStatistics.ResourcesError / projectStatistics.Pv)
		projectStatistics.HttpError = Decimal(projectStatistics.HttpError / projectStatistics.Pv)
		projectStatisticsList = append(projectStatisticsList, projectStatistics)
	}
	return
	// js 报错
	// http 报错
	// 资源报错
}

func SqlWhereBuild(model string) string {
	return "from_unixtime(" + model + ".happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ? "
}
