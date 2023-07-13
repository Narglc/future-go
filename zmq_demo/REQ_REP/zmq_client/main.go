package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	// 获取 zeromq 上下文对象
	zctx, _ := zmq.NewContext()
	// 获取 套接字对象
	s, _ := zctx.NewSocket(zmq.REQ)
	// 连接远程服务
	s.Connect("tcp://127.0.0.1:5001")

	// 向服务端发送数据
	s.Send("Hello", 0)
	// 从服务端接受响应
	msg, _ := s.Recv(0)
	fmt.Printf("Data: %s\n", msg)

}
