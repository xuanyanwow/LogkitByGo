package apis

import (
	"SiamLogKit/app/model/siam_api_log"
	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/shopspring/decimal"
)

func Overview(projectId int, date int) interface{} {

	model := siam_api_log.Model

	filed := "DATE_FORMAT(`create_time`, '%H:%i') AS time, count(id) as num, sum(CASE WHEN is_success = 1 THEN 1 ELSE 0 END) AS success_times, sum(CASE WHEN is_success = 0 THEN 1 ELSE 0 END) AS fail_times"

	res, err := model.M.Where("create_date", date).Where("project_id", projectId).Group("time").Order("time").Fields(filed).All()

	if err != nil {
		return nil
	}
	return res
}

func GetQpsInfo(second int, projectId int) interface{} {
	startTimeStamp := int(gtime.Timestamp()) - second
	startTime := gtime.NewFromTimeStamp(int64(startTimeStamp)).Format("Y-m-d H:i:s")
	endTime := gtime.Now().Format("Y-m-d H:i:s")

	model := siam_api_log.Model
	res, err := model.Where(g.Map{
		"create_time between ? and ?": g.Slice{startTime, endTime},
		"project_id":                  projectId,
	}).Count()
	if err != nil {
		return g.Map{
			"count": 0,
			"qps":   0,
		}
	}

	return g.Map{
		"count": res,
		"qps":   float64(res) / float64(second),
	}
}

func UserFromList(projectId int) interface{} {
	model := siam_api_log.Model

	res, _ := model.M.Where("project_id", projectId).Fields("user_from").Group("user_from").All()
	return res

}

func Proportion(projectId int, createDate interface{}) interface{} {
	model := siam_api_log.Model

	res, err := model.M.Where("project_id", projectId).Where("create_date", createDate).Fields(`api_full,count(id) as num,
        sum(CASE WHEN is_success = 1 THEN 1 ELSE 0 END) AS success_times,
        sum(CASE WHEN is_success = 0 THEN 1 ELSE 0 END) AS fail_times,
        avg(consume_time) as avg_consume_time`).Group("api_full").Order("num DESC").FindAll()

	if err != nil {
		return g.Map{}
	}
	totalNum := 0
	for _, value := range res {
		tem := value.Map()
		totalNum += int(tem["num"].(int64))
	}
	totalNumDecimal := decimal.New(int64(totalNum), 0)
	decimal.DivisionPrecision = 3
	for key, value := range res {
		tem := value.Map()
		temValueNum := tem["num"].(int64)
		res[key]["proportion"] = gvar.New(decimal.NewFromFloat(float64(temValueNum)).Div(totalNumDecimal).Mul(decimal.New(int64(100), 0)))

		if tem["success_times"] == 0 {
			res[key]["can_use"] = gvar.New(0)
		} else {
			res[key]["can_use"] = gvar.New(decimal.New(int64(tem["success_times"].(float64)), 0).Div(decimal.New(temValueNum, 0)).Mul(decimal.New(int64(100), 0)))
		}
	}

	return res
}

func Detail(projectId int, userFrom, userIdentify string) interface{} {
	model := siam_api_log.Model

	res, _ := model.Where(g.Map{
		"project_id":    projectId,
		"user_from":     userFrom,
		"user_identify": userIdentify,
	}).Order("id DESC").All()

	return res
}
