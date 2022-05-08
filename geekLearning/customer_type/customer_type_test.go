package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConvertionFn func(op int) int

func timeSpend(inner IntConvertionFn) IntConvertionFn {
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
	my_func := timeSpend(slowFun)
	t.Log(my_func(13))
}
