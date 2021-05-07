package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetSurveyStatisticsData(context *gin.Context) {
	var suveyParams request.SuveyParams
	err := context.BindQuery(&suveyParams)
	startTime, endTime := getTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyStatisticsData(startTime, endTime, suveyParams.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}

func GetSurveyPUvData(context *gin.Context) {
	var suveyParams request.SuveyParams
	err := context.BindQuery(&suveyParams)
	startTime, endTime := getTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyPUvData(startTime, endTime, suveyParams.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}

func GetSurveyJsErrorData(context *gin.Context) {
	var suveyParams request.SuveyParams
	err := context.BindQuery(&suveyParams)
	startTime, endTime := getTodayStartAndEndTime()
	surveyResponse, err := services.GetSurveyJsErrorData(startTime, endTime, suveyParams.MonitorId)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(surveyResponse, "获取成功", context)
	}
}
