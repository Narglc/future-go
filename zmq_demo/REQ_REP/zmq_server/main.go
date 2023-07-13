package main

import (
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// 获取 zeromq 上下文对象
	zctx, _ := zmq.NewContext()
	// 获取 套接字对象
	s, _ := zctx.NewSocket(zmq.REP)
	// 监听 5000 端口
	s.Bind("tcp://*:5001")

	for {
		// 接受客户端发送的数据
		msg, _ := s.Recv(0)
		log.Printf("Data: %s\n", msg)

		// 返回给客户端数据 Hello World
		s.Send("Hello World", 0)
	}
}
