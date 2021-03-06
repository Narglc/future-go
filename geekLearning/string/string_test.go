package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认零值”“

	s = "hello"
	t.Logf("len:%d", len(s))
	//s[1] = '3'	 	// string 是不可变的 byte slice

	s = "\xe4\xb8\xad" //"\xe4\xb8\xa5" // 可以存储任何二进制数据
	t.Logf("str:%s,len:%d\n", s, len(s))

	s = "中"
	t.Logf("中 len%d.", len(s))

	c := []rune(s)
	t.Logf("rune('zhong')=%x,len=%d.", c, len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
