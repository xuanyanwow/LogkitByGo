package logs

import (
	"SiamLogKit/app/model/siam_logs"
	"SiamLogKit/library/di"
	"github.com/gogf/gf/container/gqueue"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

func Report(data interface{}) string {
	// 校验输入参数
	reportData := siam_logs.Entity{}
	toErr := gconv.Struct(data, &reportData)
	reportData.CreateAt = gtime.Now()

	if toErr != nil {
		return toErr.Error()
	}

	if !checkRule(reportData) {
		return "检验参数失败"
	}

	// 入队列
	logsQueue := di.Get("logsQueue").(*gqueue.Queue)
	logsQueue.Push(reportData)
	return ""
}
func Reports(data interface{}) {
	// 校验输入参数

	// 入队列
}

func checkRule(obj siam_logs.Entity) bool {
	rules := map[string]string{
		"ProjectId":   "required",
		"LogCategory": "required",
		"LogPoint":    "required",
		"LogSn":       "required",
		"LogData":     "required",
		"LogFrom":     "required",
	}

	if e := gvalid.CheckStruct(obj, rules); e != nil {
		return false
	}
	return true
}
