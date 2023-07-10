package main

import (
	"fmt"
	contextdemo "future-go/basic_demo/context_demo"
	flagdemo "future-go/basic_demo/flag_demo"
	_ "future-go/basic_demo/init_demo"       // 匿名引入包，只为在main之前调用init_demo中的init函数进行资源初始化
	demo "future-go/basic_demo/package_demo" // 包的路径: 以goPath的src为根目录定义,仅到目标目录为止,且必须双引号
	"log"
)

func main() {
	fmt.Printf("hello world\n")
	log.Printf("rst:%d\n", demo.Add(13, 15))
	fmt.Printf("Mode:%d", demo.Mode)
	demo.Test()

	/* ./main.exe -u narglc -p 9527123 -h 192.254.1.16 -port 1024 */
	flagdemo.InitFlag()

	/* context demo */
	contextdemo.TestContext()
}
