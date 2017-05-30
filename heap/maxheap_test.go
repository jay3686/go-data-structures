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

	if !reflect.DeepEqual(m.data, []int{6, 4, 5, 1, 3, 2}) {
		t.Error("Multiple element insert fail: ", m.data)
	}
}
func TestDeleteOne(t *testing.T) {
	m := MaxHeap{}
	m.Insert(1)
	v := m.ExtractMax()
	if v != 1 {
		t.Error("Single element extract fail")
	}
}

func TestDeleteManyDesc(t *testing.T) {
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

func TestDeleteManyAsc(t *testing.T) {
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
