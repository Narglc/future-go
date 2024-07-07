package timeoutctx

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 参考文章
// 一文搞懂如何实现 Go 超时控制
// https://juejin.cn/post/6944855637981397029

func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	done := make(chan error)

	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// 运行: go test timeoutctx_test.go -v
func TestMain(t *testing.T) {
	const total = 1000
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

	fmt.Println("elapsed:", time.Since(now))
}
