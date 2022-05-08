package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e1 := Employee{"0", "Bob", 20}

	e2 := Employee{Name: "Mike", Age: 32}

	e3 := new(Employee) // 注意: 这里返回引用/指针，相当于 e3:=&Employee{}
	e3.Id = "2"         // 通过实例的指针访问成员不需要使用->
	e3.Age = 22
	e3.Name = "Jack"

	t.Log(e1)
	t.Log(e2)
	t.Log(e2.Id)
	t.Log(e3)
	t.Logf("e1 is %T", e1)
	t.Logf("e2 is %T", e3)
}

/*
// 第一种定义方式: 在实例对应方法被调用时，实例的成员会进行复制
func (e Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
*/

// 通常情况下，为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) String() string {
	fmt.Printf("Address is %x.\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestStructOperations(t *testing.T) {
	e := Employee{"9527", "Narglc", 32}
	fmt.Printf("Address is %x.\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
