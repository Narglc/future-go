package main

import (
	"fmt"
	"log"

	"github.com/pebbe/zmq4"
)

func main() {
	subscriber, err := zmq4.NewSocket(zmq4.SUB)

	if err != nil {
		log.Fatalf("zmq4.NewContext fail, err:%v", err)
	}

	defer subscriber.Close()

	err = subscriber.Connect("tcp://127.0.0.1:5002") // 连接到Pub（发布）端的地址

	if err != nil {
		log.Fatalf("subscriber.Connect fail, err:%v", err)
	}

	err = subscriber.SetSubscribe("") // 订阅所有消息

	if err != nil {
		log.Fatalf("subscriber.SetSubscribe fail, err:%v", err)
	}

	for {
		msg, _ := subscriber.Recv(0)
		fmt.Println("Received:", msg)
	}
}
