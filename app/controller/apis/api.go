package apis

import (
	"SiamLogKit/app/model/siam_logs"
	"SiamLogKit/app/service/apis"
	"SiamLogKit/library/response"
	"github.com/gogf/gf/container/garray"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"strconv"
)

type Controller struct{}

func (c *Controller) Overview(r *ghttp.Request) {

	dateArray := initDateArray()
	dateInt, _ := strconv.Atoi(gtime.Now().Format("Ymd"))

	projectId, _ := strconv.Atoi(r.Get("project_id").(string))
	data := apis.Overview(projectId, dateInt)
	qps := apis.GetQpsInfo(5, projectId)

	response.JsonExit(r, 200, "ok", map[string]interface{}{
		"date": dateArray,
		"data": data,
		"qps":  qps,
	})

}

func (c *Controller) UserFromList(r *ghttp.Request) {

	projectId, _ := strconv.Atoi(r.Get("project_id").(string))

	res := apis.UserFromList(projectId)

	response.JsonExit(r, 200, "ok", g.Map{
		"list":  res,
		"count": 0,
	})
}

func (c *Controller) Proportion(r *ghttp.Request) {
	projectId, _ := strconv.Atoi(r.Get("project_id").(string))

	// 可以改成筛选日期
	res := apis.Proportion(projectId, gtime.Now().Format("Ymd"))
	response.JsonExit(r, 200, "ok", g.Map{
		"list":  res,
		"count": 0,
	})
}

func initDateArray() interface{} {

	returnArr := garray.NewStrArray(true)

	h := 0
	m := 0

	add := 1
	for {
		m += add
		if m == 60 {
			m = 0
			h += 1
		}

		hStr := ""
		mStr := ""
		if h < 10 {
			hStr = "0" + strconv.Itoa(h)
		} else {
			hStr = strconv.Itoa(h)
		}
		if m < 10 {
			mStr = "0" + strconv.Itoa(m)
		} else {
			mStr = strconv.Itoa(m)
		}

		tem := hStr + ":" + mStr
		returnArr.Append(tem)
		if h == 24 {
			return returnArr
		}
	}

}

func (c *Controller) Detail(r *ghttp.Request) {
	projectId, projectIdErr := strconv.Atoi(r.Get("project_id").(string))
	if projectIdErr != nil {
		response.JsonExit(r, 500, "error", projectIdErr.Error())
	}
	userFrom := r.Get("user_from").(string)
	userIdentify := r.Get("user_identify").(string)

	res := apis.Detail(projectId, userFrom, userIdentify)

	if res == nil {
		res = []*siam_logs.Entity{}
	}

	response.JsonExit(r, 200, "success", res)
}
