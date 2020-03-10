package index

import (
	"SiamLogKit/app/service/install"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct{}

func (c *Controller) Index(r *ghttp.Request) {
	// 判断是否创建了表结构
	if !install.IsInstall() {
		r.Response.RedirectTo("/install")
	}
	// 显示后台
	r.Response.Write("test")
}
