package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	log.Printf("----> got here\n")
	sub, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatalf("NewSocket err:%v", err)
	}
	defer sub.Close()

	// 连接到发布端的地址和端口
	err = sub.Connect("tcp://127.0.0.1:5002")
	if err != nil {
		log.Fatalf("Connect err:%v\n", err)
	}

	// 订阅所有消息
	err = sub.SetSubscribe("")
	if err != nil {
		log.Fatalf("SetSubscribe err:%v\n", err)
	}
	for i := 0; i < 10; i++ {
		// 接收订阅端的消息
		log.Printf("Prepare to recv msg...\n")
		msg, err := sub.Recv(0)
		if err != nil {
			log.Printf("Recv fail, err:%v\n", err)
			continue
		}
		fmt.Println("收到消息:", msg)
	}
}
