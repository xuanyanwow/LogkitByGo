package boot

import (
	"SiamLogKit/app/consume"
	"SiamLogKit/library/di"
	"github.com/gogf/gf/container/gqueue"
)

// 框架需要的队列初始化配置
func init() {
	logsQueue := gqueue.New()
	di.Set("logsQueue", logsQueue)
	// 启动监听
	go consume.Logs(logsQueue)

	apiQueue := gqueue.New()
	di.Set("apiQueue", apiQueue)
	// 启动监听
	go consume.Apis(apiQueue)
}
