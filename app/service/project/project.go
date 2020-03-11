package project

import (
	"SiamLogKit/app/model/siam_projects"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)



func GetList() *gmap.AnyAnyMap{
	data,err := siam_projects.Model.All()

	returnData := gmap.New()
	if err != nil {
		return returnData
	}
	count := len(data)

	returnData.Set("list", data)
	returnData.Set("count", count)

	return returnData
}


func Add(projectName string) bool {
	res,err := siam_projects.Model.Data(g.Map{
		"project_name" : projectName,
	}).Save()

	if err != nil {
		return false
	}

	rowNum, _ := res.RowsAffected()
	return rowNum > 0
}

func DeleteOne(projectId int) bool {
	res, err := siam_projects.Model.Delete("project_id", projectId)
	if err != nil {
		return false
	}

	rowNum, _ := res.RowsAffected()
	return rowNum > 0
}