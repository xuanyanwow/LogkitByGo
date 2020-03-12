package apis

import (
	"SiamLogKit/app/model/siam_api_log"
	"SiamLogKit/library/di"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
	"strconv"
)

func Reports(data []interface{}) string {
	for _, value := range data {
		tem := Report(value)
		if tem != "" {
			return tem
		}
	}

	return ""
}

func Report(data interface{}) string {

	// 校验输入参数
	reportData := siam_api_log.Entity{}
	toErr := gconv.Struct(data, &reportData)

	reportData.CreateDate, _ = strconv.Atoi(gtime.Now().Format("Ymd"))
	reportData.CreateTime = gtime.Now()
	reportData.ApiFull = reportData.ApiCategory + "/" + reportData.ApiMethod

	if toErr != nil {
		return toErr.Error()
	}

	if !checkRule(reportData) {
		return "检验参数失败"
	}

	// 入队列
	logsQueue := di.Get("apiQueue").(*gqueue.Queue)
	logsQueue.Push(reportData)
	return ""
}

func checkRule(data siam_api_log.Entity) bool {
	rules := map[string]string{
		"ProjectId":    "required",
		"ApiCategory":  "required",
		"ApiMethod":    "required",
		"IsSuccess":    "required",
		"ConsumeTime":  "required",
		"UserFrom":     "required",
		"UserIdentify": "required",
	}

	if e := gvalid.CheckStruct(data, rules); e != nil {
		return false
	}
	return true
}
