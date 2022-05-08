package condition

import (
	"fmt"
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	/*
		if v, err := someFun(); err == nil {
			// todo
		} else {
			// todo2
		}
	*/
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
