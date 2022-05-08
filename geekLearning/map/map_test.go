package my_map

import "testing"

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 4, 3: 9}
	t.Logf("%v m1[2]=%d,m1 len(%d)", m1, m1[2], len(m1)) // map[1:2 2:4 3:9] m1[2]=4,len(3)

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2 len(%d)", len(m2)) // m2 len(1)

	m3 := make(map[int]int, 10)
	t.Logf("m3 len(%d)", len(m3)) // m3 len(0)
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 0

	m1[2] = 0
	t.Log(m1[2]) // 0
	t.Log(m1)    // map[2:0]

	if val, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", val)
	} else {
		t.Log("Key 3 is not existing.") // Key 3 is not existing.
	}
}

func TestMapTravel(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
