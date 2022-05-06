package main

import (
	"fmt"
	_ "future-go/basic_demo/init_demo"       // 匿名引入包，只为在main之前调用init_demo中的init函数进行资源初始化
	demo "future-go/basic_demo/package_demo" // 包的路径: 以goPath的src为根目录定义,仅到目标目录为止,且必须双引号
	"log"
)

func main() {
	fmt.Printf("hello world\n")
	log.Printf("rst:%d\n", demo.Add(13, 15))
	fmt.Printf("Mode:%d", demo.Mode)
	demo.Test()
}
