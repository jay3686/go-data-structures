package heap

import (
	"reflect"
	"testing"
)

func TestMinInsertOne(t *testing.T) {
	m := MinHeap{}
	m.Insert(1)
	if m.data[0] != 1 {
		t.Error("Single element insert fail")
	}
}

func TestMinInsertManyDesc(t *testing.T) {
	m := MinHeap{}
	nums := []int{6, 5, 4, 3, 2, 1}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{1, 3, 2, 6, 4, 5}) {
		t.Error("Multiple element insert fail", m.data)
	}
}

func TestMinInsertManyAsc(t *testing.T) {
	m := MinHeap{}
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Multiple element insert fail: ", m.data)
	}
}

func TestMinDeleteOne(t *testing.T) {
	m := MinHeap{}
	m.Insert(1)
	v := m.ExtractMin()
	if v != 1 {
		t.Error("Single element extract fail")
	}
}

func TestMinDeleteFromEmptyHeap(t *testing.T) {
	m := MinHeap{}
	v := m.ExtractMin()
	if v != 0 {
		t.Error("Empty extract fail")
	}
}

func TestMinDeleteManyDesc(t *testing.T) {
	m := MinHeap{}
	nums := []int{6, 5, 4, 3, 2, 1}

	for _, n := range nums {
		m.Insert(n)
	}

	minedNums := []int{1, 2, 3, 4, 5, 6}
	for _, n := range minedNums {
		v := m.ExtractMin()
		if v != n {
			t.Error("Multi element extract fail", v, m.data)
		}
	}
}

func TestMinDeleteManyAsc(t *testing.T) {
	m := MinHeap{}
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		m.Insert(n)
	}

	for _, n := range nums {
		v := m.ExtractMin()
		if v != n {
			t.Error("Multi element extract fail", v, m.data)
		}
	}
}

func TestMinPeekOne(t *testing.T) {
	m := MinHeap{}
	m.Insert(1)
	v := m.FindMin()
	if v != 1 {
		t.Error("Single element peek fail")
	}
}

func TestMinPeekFromEmptyHeap(t *testing.T) {
	m := MinHeap{}
	v := m.FindMin()
	if v != 0 {
		t.Error("Empty peek fail")
	}
}

func TestMinLen(t *testing.T) {
	m := MinHeap{}
	if m.Len() != 0 {
		t.Error("empty len is fail")
	}
	m.Insert(1)
	m.Insert(2)
	m.Insert(3)
	m.Insert(4)
	if m.Len() != 4 {
		t.Error("len is fail")
	}
}
