package project

import (
	"SiamLogKit/app/service/project"
	"SiamLogKit/library/response"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

type Controller struct{}


func (c *Controller) GetList(r *ghttp.Request) {
	data := project.GetList()
	response.JsonExit(r, 200, "success", data)
}

func (c *Controller) Add(r *ghttp.Request) {
	projectName := r.Get("project_name").(string)
	res := project.Add(projectName)
	response.JsonExit(r, 200, "success", res)
}

func (c *Controller) DeleteOne(r *ghttp.Request) {
	projectId,err := strconv.Atoi(r.Get("project_id").(string))
	if err != nil {
		response.JsonExit(r, 500, "error")
	}

	res := project.DeleteOne(projectId)
	response.JsonExit(r, 200, "success", res)
}