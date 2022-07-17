package series

import "fmt"

func GetFibonacci(n int) []int {
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList
}

func square(n int) int {
	return n * n
}

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}
