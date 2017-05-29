package heap

import "testing"

func TestInsert(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	if m.data[0] != 1 {
		t.Error("Single element insert fail")
	}
}
