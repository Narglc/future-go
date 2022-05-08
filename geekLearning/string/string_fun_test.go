package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	t.Log(strings.Join(parts, "_"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Logf("str:" + s)

	if rst, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + rst)
	}

}
