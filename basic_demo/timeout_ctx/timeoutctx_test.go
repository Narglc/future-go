package timeoutctx

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// 参考文章
// 一文搞懂如何实现 Go 超时控制
// https://juejin.cn/post/6944855637981397029

func hardWork(job interface{}) error {
	fmt.Println("start hardwork...")
	time.Sleep(10 * time.Second)
	fmt.Println("end hardwork...")
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan error, 2) // 此处设置缓冲/非缓冲差别很大： 缓冲无goroutine泄漏，非缓冲有goroutine泄漏

	go func() {
		fmt.Println("before go hardwork...")
		done <- hardWork(job)
		fmt.Println("after go hardwork...")
	}()

	select {
	case err := <-done:
		fmt.Printf("got here...\n")
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 运行: go test timeoutctx_test.go -v
func TestMain(t *testing.T) {
	const total = 10
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()

	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "just test...")
		}()
	}
	wg.Wait()

	fmt.Println("elapsed:", time.Since(now)) // elapsed: 2.005725931s
	time.Sleep(time.Second * 20)
	fmt.Println("number of goroutines:", runtime.NumGoroutine()) // number of goroutines: 1002
}
