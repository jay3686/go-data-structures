package heap

import "testing"

import "reflect"

func TestInsertOne(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	if m.data[0] != 1 {
		t.Error("Single element insert fail")
	}
}

func TestInsertManyDesc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{6, 5, 4, 3, 2, 1}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{6, 5, 4, 3, 2, 1}) {
		t.Error("Multiple element insert fail")
	}
}

func TestInsertManyAsc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{6, 5, 4, 3, 2, 1}) {
		t.Error("Multiple element insert fail")
	}
}
