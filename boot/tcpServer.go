package boot

import (
	"SiamLogKit/app/tcp"
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"io"
)

func init() {
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		// 触发onOpen
		fmt.Println(tcp.Encode("Test"))
		conn.SendPkg([]byte("siam你好"))
		for {
			data, err := conn.RecvPkg()
			if err != nil {
				if err == io.EOF {
					// 触发onClose
					fmt.Println("onclose")
				}else{
					fmt.Println(err)
				}
				break
			}
			// 触发onMessage
			fmt.Println("receive:", string(data))
		}
	}).Run()
}

