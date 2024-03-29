package err_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThenHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}

	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error:", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(list)
}

func TestGetFibonacci(t *testing.T) {
	if rst, err := GetFibonacci(1); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		}
		t.Error(err)
	} else {
		t.Log(rst)
	}

	GetFibonacci2("13")
}

func TestNarglc(t *testing.T) {
	//TODO
}
