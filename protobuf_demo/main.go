package main

import (
	"fmt"
	"future-go/protobuf_demo/proto/hello"

	"google.golang.org/protobuf/proto"
)

func main() {
	pb2Send := &hello.Say{
		Id:    9527,
		Hello: "say hello",
		Word:  []string{"堕落的世界", "伤心太平洋"},
	}

	//用字符串的方式：打印ProtoBuf消息
	fmt.Printf("字符串输出结果：%v\n", pb2Send.String())

	//转成二进制文件
	marshal, err := proto.Marshal(pb2Send)
	if err != nil {
		return
	}
	fmt.Printf("Marshal转成二进制结果：%v\n", marshal)

	//将二进制文件转成结构体
	newpb2Send := hello.Say{}
	err = proto.Unmarshal(marshal, &newpb2Send)
	if err != nil {
		return
	}
	fmt.Printf("二进制转成结构体的结果：%v\n", &newpb2Send)
}
