package boot

import (
	"fmt"
	"github.com/gogf/gf/net/gtcp"
	"io"
)

func init() {
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		// 触发onOpen
		for {
			data, err := conn.Recv()
			if err != nil {
				if err == io.EOF {
					fmt.Println("onclose")
				}else{
					fmt.Println(err)
				}
				break
			}
			// 触发onMessage
			fmt.Println("receive:", data)
		}
	}).Run()
}

