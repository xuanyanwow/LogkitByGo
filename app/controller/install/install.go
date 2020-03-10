package install

import (
	"SiamLogKit/app/service/install"
	"SiamLogKit/app/service/mysql"
	"SiamLogKit/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gins"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"io/ioutil"
)

type Controller struct {
}

func (c *Controller) Index(r *ghttp.Request) {
	// 判断是否创建了表结构
	if install.IsInstall() {
		r.Response.Write("已经安装过了，无需重复安装")
		return
	}

	versionInfo, err := mysql.GetVersion()
	if err != nil {
		response.Echo(r, err.Error()+"<br/>")
		response.Echo(r, "读取mysql版本失败，请检查数据库配置")
		return
	}
	if len(versionInfo) <= 0 {
		response.Echo(r, "读取mysql版本失败，请重试")
		return
	}

	mapData := versionInfo.Map()
	databaseInfo := gins.Config().GetMap("database")["default"].(g.Array)

	r.Response.WriteTpl("/install/install.tpl", g.Map{
		"version":      mapData["version()"],
		"databaseInfo": databaseInfo[0],
	})
}

func (c *Controller) Run(r *ghttp.Request) {
	// 表不存在则创建表
	if mysql.IsExistTable("siam_logs") {
		response.EchoExit(r, "表结构已存在")
	}
	if !mysql.CreateTable("siam_logs") {
		response.EchoExit(r, "创建失败")
	}

	// 生成锁
	ioutil.WriteFile("./public/install.lock", []byte(gtime.Datetime()), 0777)

	response.EchoExit(r, "安装成功 <a href='/'>返回</a>")
}
