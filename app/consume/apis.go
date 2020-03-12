package consume

import (
	"SiamLogKit/app/model/siam_api_log"
	"fmt"
	"github.com/gogf/gf/container/gqueue"
	"time"
)

func Apis(queue *gqueue.Queue) {
	for {
		data := queue.Pop()

		// 插入数据库
		res, err := siam_api_log.Model.Insert(data)
		fmt.Println(res, err)

		time.Sleep(time.Millisecond)
	}
}
