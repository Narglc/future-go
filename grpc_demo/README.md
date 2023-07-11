# 开发gRPC的流程

- 写proto文件定义服务和消息
- 使用protoc工具生成代码
- 编写业务逻辑代码提供服务

在微服务架构开发gRPC时，一定有两个端：服务端和客户端。

我们的习惯是，在搞定protobuf之后，先写服务端逻辑，暴露端口，提供服务；再写客户端逻辑，连接服务，发送请求，处理响应。


## 安装protobuf生成插件

```bash
# 安装protobuf
brew install ProtoBuf

# 安装 protoc-gen-go
go get github.com/golang/protobuf/protoc-gen-go

# 安装 grpc 相关插件
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 编译protobuf

```bash
# 如果which protoc-gen-go 找不到，则需要将$GOBIN 路径加入PATH
export PATH=$PATH:/Users/narglc/go/bin

# 指定生成pb.go,grpc.pb.go
protoc --go_out=. --go-grpc_out=. helloworld.proto

# 或者使用 --plugin 指定 protoc-gen-go 的路径
protoc --plugin=../../../bin/protoc-gen-go --go_out=. proto/hello.proto
```

## 运行

```bash
# 服务端
go run main.go --port=9527

# 客户端
go run main.go --addr=localhost:9527 --name=Narglc
```