package contextdemo

import (
	"context"
	"fmt"
	"time"
)

func value() {
	// 初识ctx使用Backgroud、TODO都可以
	ctx := context.Background() // context.TODO()

	// 添加key-value键值对
	ctx = context.WithValue(ctx, "Name", "Narglc")
	ctx = context.WithValue(ctx, "Age", 18)
	ctx = context.WithValue(ctx, "Addr", "Shenzhen Nanshan Xili")

	keys := []string{"Name", "Age", "Addr"}

	for idx, key := range keys {
		fmt.Println(idx, ":", key, " --> ", ctx.Value(key))
	}
}

// context.WithTimeout, ctx.Done(), ctx.Err()
func timeout() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*1) //time.Millisecond * 1000
	t0 := time.Now()
	defer cancel()

	select {
	case <-ctx.Done(): // ctx 被取消时，则返回

		// 使用 time.Since 计算耗时
		fmt.Println("time.Since:", time.Since(t0).Milliseconds())
		err := ctx.Err()
		fmt.Println("ctx err:", err)
	}
}

func cancel() {
	ctx, cancel := context.WithCancel(context.TODO())
	t0 := time.Now()

	go func() {
		time.Sleep(time.Millisecond * 100)
		cancel()
	}()

	select {
	case <-ctx.Done():
		t1 := time.Now()
		// 使用 time.Sub 计算耗时
		fmt.Println("time.Sub:", t1.Sub(t0).Milliseconds())
		err := ctx.Err()
		fmt.Println(err)
	}
}

func TestContext() {
	// 计算函数耗时
	t0 := time.Now()
	defer func() {
		fmt.Printf("cur func Cost %d ms\n", time.Since(t0).Milliseconds())
	}()

	fmt.Println("TestContext:hello world  ==============")

	value()
	timeout()
	cancel()
}
