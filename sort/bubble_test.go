package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	BubbleSort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Not sorted in asc order", nums)
	}
}
func TestBubbleSort2(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	BubbleSort2(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Not sorted in asc order", nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < b.N; i++ {
		BubbleSort2(nums)
	}
}
