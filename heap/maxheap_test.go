package heap

import (
	"reflect"
	"testing"
)

func TestMaxInsertOne(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	if m.data[0] != 1 {
		t.Error("Single element insert fail")
	}
}

func TestMaxInsertManyDesc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{6, 5, 4, 3, 2, 1}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{6, 5, 4, 3, 2, 1}) {
		t.Error("Multiple element insert fail")
	}
}

func TestMaxInsertManyAsc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		m.Insert(n)
	}

	if !reflect.DeepEqual(m.data, []int{6, 4, 5, 1, 3, 2}) {
		t.Error("Multiple element insert fail: ", m.data)
	}
}

func TestMaxDeleteOne(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	v := m.ExtractMax()
	if v != 1 {
		t.Error("Single element extract fail")
	}
}

func TestMaxDeleteFromEmptyHeap(t *testing.T) {
	m := MaxHeap{}
	v := m.ExtractMax()
	if v != 0 {
		t.Error("Empty extract fail")
	}
}

func TestMaxDeleteManyDesc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{6, 5, 4, 3, 2, 1}

	for _, n := range nums {
		m.Insert(n)
	}

	for _, n := range nums {
		v := m.ExtractMax()
		if v != n {
			t.Error("Multi element extract fail", v, m.data)
		}
	}

}

func TestMaxDeleteManyAsc(t *testing.T) {
	m := MaxHeap{}
	nums := []int{1, 2, 3, 4, 5, 6}

	for _, n := range nums {
		m.Insert(n)
	}

	maxedNums := []int{6, 5, 4, 3, 2, 1}
	for _, n := range maxedNums {
		v := m.ExtractMax()
		if v != n {
			t.Error("Multi element extract fail", v, m.data)
		}
	}
}

func TestMaxPeekOne(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	v := m.FindMax()
	if v != 1 {
		t.Error("Single element peek fail")
	}
}

func TestMaxPeekFromEmptyHeap(t *testing.T) {
	m := MaxHeap{}
	v := m.FindMax()
	if v != 0 {
		t.Error("Empty peek fail")
	}
}

func TestMaxLen(t *testing.T) {
	m := MaxHeap{}
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
