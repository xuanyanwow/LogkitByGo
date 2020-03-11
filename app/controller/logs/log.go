package logs

import (
	"SiamLogKit/app/model/siam_logs"
	"SiamLogKit/app/service/logs"
	"SiamLogKit/library/response"
	"github.com/gogf/gf/net/ghttp"
	"strconv"
)

type Controller struct{}



func (c *Controller) Query(r *ghttp.Request) {
	projectId,projectIdErr := strconv.Atoi(r.Get("project_id").(string))
	if projectIdErr != nil {
		response.JsonExit(r, 500, "error", projectIdErr.Error())
	}
	logSn     := r.Get("log_sn").(string)


	res := logs.QueryByProjectSn(projectId, logSn)

	if res == nil {
		res = []*siam_logs.Entity{}
	}

	response.JsonExit(r, 200, "success", res)
}
