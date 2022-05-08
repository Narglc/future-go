package my_map

import "testing"

func TestMapWithValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true

	n := 1
	if mySet[n] {
		t.Log("exist")
	} else {
		t.Log("not exist")
	}

	mySet[3] = true
	t.Log(len(mySet))

	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Log("exist")
	} else {
		t.Log("not exist")
	}

}
