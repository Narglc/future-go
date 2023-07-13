package main

import (
	"fmt"
	"log"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	pub, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		log.Fatalf("NewSocket err:%v", err)
	}
	defer pub.Close()

	// 绑定发布端套接字到指定IP和端口
	err = pub.Bind("tcp://127.0.0.1:5002")
	if err != nil {
		log.Fatalf("Bind err:%v\n", err)
	}
	// 要给予sub connect的时间，三次握手完成后再进行Pub
	time.Sleep(time.Second * 1)

	for i := 0; ; i++ {
		msg := fmt.Sprintf("消息 %d", i)

		// 发布消息
		n, err := pub.Send(msg, 0)
		if err != nil {
			log.Fatalf("Send err:%v\n", err)
		}
		fmt.Printf("已发布消息(%d):%s\n", n, msg)

		if i == 9 {
			break
		}
	}
	time.Sleep(time.Second * 10)
}
