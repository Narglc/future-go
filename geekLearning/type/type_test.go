package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64 = 2
	b = int64(a)

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1  // 不支持指针运算
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Logf("[%s]", s)
	t.Log(len(s))
}
