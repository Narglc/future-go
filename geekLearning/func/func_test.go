package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	my_func := timeSpend(slowFun)
	t.Log(my_func(13))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(2, 3, 5, 7, 8, 9, 1))
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("Clear resources")
	}() // 匿名函数
	t.Log("Started")
	panic("Fatal error.") // defer 仍会执行
}

func Clear() {
	fmt.Println("Clear resources")
}

func TestDefer2(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
