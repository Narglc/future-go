package demo

import "fmt"

var num = 100

const Mode = 1

type person struct {
	name string
	Age  int
}

type Student struct {
	Name  string
	class string
}

func Add(x, y int) int {
	return x + y
}

func sayHi() {
	var myName = "Narglc"
	fmt.Printf("name:%s\n", myName)
	fmt.Printf("仅内部可见:%v", Student{"小明", "平中6班"})
}

func Test() {
	sayHi()
}
