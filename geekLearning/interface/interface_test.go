package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
	Work()
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}

func (g *GoProgrammer) Work() {
	fmt.Println("没日没夜地加班呀...")
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p.Work()
}
