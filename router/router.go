package router

import (
	"SiamLogKit/app/controller/apis"
	"SiamLogKit/app/controller/index"
	"SiamLogKit/app/controller/install"
	"SiamLogKit/app/controller/logs"
	"SiamLogKit/app/controller/project"
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

		s.Group("/api/", func(group *ghttp.RouterGroup) {
			ProjectController := project.Controller{}
			group.ALL("/project/get_list", ProjectController.GetList)
			group.ALL("/project/add", ProjectController.Add)
			group.ALL("/project/delete_one", ProjectController.DeleteOne)

			LogController := logs.Controller{}
			group.ALL("logs/query", LogController.Query)

			ApiController := apis.Controller{}
			group.ALL("api_log/overview", ApiController.Overview)
			group.ALL("api_log/user_from_list", ApiController.UserFromList)
			group.ALL("api_log/proportion", ApiController.Proportion)
			group.ALL("api_log/detail", ApiController.Detail)
		})
	})

}
