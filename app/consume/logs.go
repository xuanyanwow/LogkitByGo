package consume

import (
	"SiamLogKit/app/model/siam_logs"
	"fmt"
	"github.com/gogf/gf/container/gqueue"
	"time"
)

func Logs(queue *gqueue.Queue) {
	for {
		data := queue.Pop()

		// 插入数据库
		res, err := siam_logs.Model.Insert(data)
		fmt.Println(res, err)

		time.Sleep(time.Millisecond)
	}
}
