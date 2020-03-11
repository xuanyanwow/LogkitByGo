package logs

import (
	"SiamLogKit/app/model/siam_logs"
)

func QueryByProjectSn(projectId int, logSn string) []*siam_logs.Entity {
	db := siam_logs.Model
	res, err := db.Where("project_id", projectId).Where("log_sn",logSn).FindAll()
	if err != nil {
		return nil
	}
	return res
}