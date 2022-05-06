package init_demo

import (
	"fmt"
	"log"
)

var x int8 = 10

const pi = 3.14

const (
	name = "narglc"
	age  = 32
	addr = "Shenzhen"
)

func init() {
	log.Printf("------- package包匿名引入，调用init初始化 -------")
	fmt.Println("x:", x)
	fmt.Println("pi:", pi)
	fmt.Printf("name:%v, age:%v, addr:%v\n", name, age, addr)
	log.Printf("------- package包匿名引入结束 -------")
}

func init() {
	log.Println("是否会被调用呢??")
	log.Printf("是的！ init不会冲突, 都会被调用")
}
