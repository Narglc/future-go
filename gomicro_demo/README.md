# 学习笔记

## 使用cli脚手架

### 简单rpc项目
```bash
# 下载脚手架
go install github.com/go-micro/cli/cmd/go-micro@latest

# 创建项目
go-micro new service helloworld
cd helloworld
make init proto update tidy

# 编译运行
cd helloworld
make proto tidy
go-micro run

# 测试
go-micro call helloworld Helloworld.Call '{"name": "John"}'
```
**注**： 可直接通过`make proto`进行protobuf的编译。

### 包含链路追踪的rpc微服务
```bash
# 创建项目
go-micro new service --jaeger traceRpcSvr

## 如上述编译运行
```
 #### 使用docker-compose部署服务
 ```
 version: "3.7"

services:

  # -----------------------------
  # jaeger servcie
  # -----------------------------
  jaeger:
    image: jaegertracing/all-in-one:1.20
    ports:
      - "1314:6831/udp"
      - "1315:16686"
    networks:
      - backend

# -----------------------------
# networks
# -----------------------------
networks:
  backend:
    external: true
 ```