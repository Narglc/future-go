package fib

import (
	"testing"
)

func TestFirstTry(t *testing.T) {
	var (
		a int = 1
		b     = 1
	)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
