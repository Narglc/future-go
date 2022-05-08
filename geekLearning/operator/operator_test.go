package operator_test

import "testing"

const (
	Readable   = 1 << iota // 1 = 001
	Writeable              // 2 = 010
	Executable             // 4 = 100
)

func TestBitClear(t *testing.T) {
	t.Log(Readable, Writeable, Executable)

	a := 7 // 0b111
	unReadable := a &^ Readable
	t.Log(unReadable)
	unWriteable := a &^ Writeable
	t.Log(unWriteable)
	unExecutable := a &^ Executable
	t.Log(unExecutable)
}

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 2}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// t.Log(a == c)  // 长度不同的数组不可以比较
	t.Log(a == d)
}
