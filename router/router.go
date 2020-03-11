package router

import (
    "SiamLogKit/app/controller/index"
	"SiamLogKit/app/controller/install"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		IndexController := index.Controller{}
		group.ALL("/", IndexController.Index)
		InstallController := install.Controller{}
		group.ALL("/install", InstallController.Index)
		group.ALL("/install_run_sql", InstallController.Run)
	})
}
