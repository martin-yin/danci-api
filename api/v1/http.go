package v1

import (
	"danci-api/model/request"
	"danci-api/model/response"
	"danci-api/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHttpInfo(context *gin.Context) {
	var queryPagePerformance request.QueryPagePerformance
	err := context.BindQuery(&queryPagePerformance)
	HttpInfoListResponse, err := services.GetHttpInfoList(queryPagePerformance.StartTime, queryPagePerformance.EndTime)
	HttpQuotaResponse, err := services.GetHttpQuota(queryPagePerformance.StartTime, queryPagePerformance.EndTime)
	if err != nil {
		//global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PageHttpResponse{
			HttpListResponse:  HttpInfoListResponse,
			HttpQuotaResponse: HttpQuotaResponse,
		}, "获取成功", context)
	}
}

func GetHttpStage(context *gin.Context) {
	var queryPagePerformance request.QueryPagePerformance
	err := context.BindQuery(&queryPagePerformance)
	HttpStageTimeResponse, err := services.GetHttpStageTime(queryPagePerformance.StartTime, queryPagePerformance.EndTime, queryPagePerformance.TimeGrain, queryPagePerformance.StageType)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取失败：%v", err), context)
	} else {
		response.OkWithDetailed(response.PageHttpStage{
			HttpStageTimeResponse: HttpStageTimeResponse,
		}, "获取成功", context)
	}
}
