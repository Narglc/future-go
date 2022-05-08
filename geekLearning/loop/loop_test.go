package loop

import (
	"fmt"
	"testing"
)

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}

	for n := 0; n < 9; n++ {
		fmt.Println(n)
	}
}
