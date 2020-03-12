package boot

import (
	"SiamLogKit/app/tcp/logs"
	"SiamLogKit/library/tcp"
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"io"
)

func init() {
	tcpServer := tcp.Tcp{}
	// 设置路由映射
	tcpServer.SetFunctionMap(map[string]interface{}{
		"report":    logs.Report,
		"reports":   logs.Reports,
		"api":       nil,
		"apis":      nil,
		"exception": nil,
	})

	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		// 触发onOpen
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				if err == io.EOF {
					// 触发onClose
				} else {
					fmt.Println(err)
				}
				break
			}
			// 触发onMessage
			// 解析执行
			tcpServer.Parse(conn, string(data))
		}
	}).Run()
}
