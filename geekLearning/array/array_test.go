package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3}
	arr2 := [...]int{1, 3, 4, 5}
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)

	c := [2][2]int{{1, 2}, {3, 4}}
	t.Log(c)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5}

	for i, v := range arr3 {
		t.Log(i, v)
	}

	c := [...][2]int{{1, 2}, {3, 4}}
	for _, v := range c {
		t.Log(v)
	}
}

func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5, 6}
	arr3_sec := arr3[:4]
	t.Log(arr3_sec)
	arr3_sec2 := arr3[:]
	t.Log(arr3_sec2)
}
