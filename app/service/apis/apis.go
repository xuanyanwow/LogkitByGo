package apis

import (
	"SiamLogKit/app/model/siam_api_log"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
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
