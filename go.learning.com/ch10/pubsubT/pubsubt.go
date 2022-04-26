package pubsubt

import (
	"fmt"
	"strings"
	"time"

	"go.learning.com/ch10/pubsub"
)

func main() {
	p := pubsub.NewPublisher(100*time.Microsecond, 10)

	defer p.Close()

	// 添加一个所有主题的订阅者（即消息队列）
	all := p.Subscribe()
	// 添加一个订阅主题为golang
	// 本质为一个有名字的通道
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}

		return false
	})

	// 发布消息
	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	// 从订阅者中拉取消息
	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	// 运行一定时间后退出
	time.Sleep(3 * time.Second)
}
