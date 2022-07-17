package client

import (
	"future-go/geekLearning/package/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(9))
	//t.Log(series.square(5)) 		// 小写函数无法外部访问
}
