package services

import (
	"danci-api/global"
	"danci-api/model"
	"danci-api/model/response"
	"fmt"
	"strconv"
)

func GetHttpInfoList(monitorId string, startTime string, endTime string) (httpInfoList []response.HttpListResponse, err error) {
	sqlWhere := `from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?`
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("http_url, Count(DISTINCT http_url) as total, Count(DISTINCT user_id) as user_total, round( AVG( load_time ), 2 ) AS load_time").Where(sqlWhere, startTimes, endTimes, monitorId).Group("http_url").Find(&httpInfoList).Debug().Error
	return
}

func GetHttpQuota(monitorId string, startTime string, endTime string) (httpQuota response.HttpQuotaResponse, err error) {
	startTimes := startTime + " 00:00:00"
	endTimes := endTime + " 23:59:59"
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("round(AVG( load_time )) AS load_time, COUNT(*) AS total").Where("page_https.`status` != 0 And from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s') and monitor_id = ?", startTimes, endTimes, monitorId).Find(&httpQuota).Error
	return
}

func GetHttpStageTimeSuccess(startTime string, endTime string) (httpStageTime []response.HttpStageTimeResponse, err error) {
	err = global.GVA_DB.Model(&model.PageHttp{}).Select("FROM_UNIXTIME( happen_time / 1000, '%H:%i') AS time_key, COUNT( * ) AS total, round( AVG( load_time ), 2 ) AS load_time").Group("time_key").Where("from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime+" 00:00:00", endTime+" 23:59:59").Find(&httpStageTime).Error
	return
}

//func GetHttpStageTimeError(startTime string, endTime string, timeGrain string) (httpStageTime []response.HttpStageTimeResponseError, err error) {
//	query := ""
//	startQuery := ""
//	endQuery := ""
//
//	if timeGrain == "minute" {
//		query = query + "'%H:%i'"
//		startQuery = "CONCAT(" + "'" + startTime + " '" + ", time_key, ':00')," + startQuery
//		endQuery = "CONCAT(" + "'" + endTime + " '" + ", time_key, ':59')," + endQuery
//	}
//	if timeGrain == "hour" {
//		query = query + "'%d %H'"
//		startQuery = "CONCAT(" + "'" + startTime[0:7] + "-'" + ", time_key, ':00:00')," + startQuery
//		endQuery = "CONCAT(" + "'" + endTime[0:7] + "-'" + ", time_key, ':59:59')," + endQuery
//	}
//	if timeGrain == "day" {
//		query = query + "'%m-%d'"
//		startQuery = "CONCAT(" + "'" + startTime[0:5] + "'" + ", time_key, '00:00:00')," + startQuery
//		endQuery = "CONCAT(" + "'" + endTime[0:5] + "'" + ", time_key, '23:59:59')," + endQuery
//	}
//	err = global.GVA_DB.Model(&model.PageHttp{}).Select(""+
//		"FROM_UNIXTIME( happen_time / 1000, "+query+") AS time_key,"+
//		"COUNT( * ) AS total, "+
//		"(SELECT COUNT( * ) FROM page_https WHERE page_https.`status`  > 305 "+
//		"AND from_unixtime( page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s' ) BETWEEN date_format("+startQuery+"'%Y-%m-%d %H:%i:%s') "+
//		"AND date_format("+endQuery+"'%Y-%m-%d %H:%i:%s')"+
//		")  AS fail_total, "+
//		"(SELECT round( AVG( load_time ) ) AS load_time from page_https WHERE page_https.`status` > 305 "+
//		"AND from_unixtime( page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s' ) BETWEEN date_format("+startQuery+"'%Y-%m-%d %H:%i:%s') "+
//		"AND date_format("+endQuery+"'%Y-%m-%d %H:%i:%s')"+
//		") AS load_time").Group("time_key").Where("from_unixtime(page_https.happen_time / 1000, '%Y-%m-%d %H:%i:%s') between date_format( ? , '%Y-%m-%d %H:%i:%s') and date_format( ?, '%Y-%m-%d %H:%i:%s')", startTime+" 00:00:00", endTime+" 23:59:59").Find(&httpStageTime).Error
//	return
//}

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
