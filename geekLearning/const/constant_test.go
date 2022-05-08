package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable   = 1 << iota // 1 = 001
	Writeable              // 2 = 010
	Executable             // 4 = 100
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry2(t *testing.T) {
	t.Log(Readable, Writeable, Executable)

	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executable == Executable)
}
