# 使用指南

## 环境准备

### 安装必备工具

```bash
brew install ProtoBuf
go get github.com/golang/protobuf/protoc-gen-go
```

### 生成go代码

```bash
cd proto/

# 如果which protoc-gen-go 找不到，则需要将$GOBIN 路径加入PATH
# 或者使用 --plugin 指定 protoc-gen-go 的路径
protoc --plugin=../../../bin/protoc-gen-go --go_out=. proto/hello.proto
```

### 编写 主函数

```golang
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
```