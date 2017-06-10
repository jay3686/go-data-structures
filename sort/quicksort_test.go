package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	QuickSort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Not sorted in asc order", nums)
	}
}

func BenchmarkQuickSortReverseOrder(b *testing.B) {
	nums := []int{}
	for i := 1000; i >= -1000; i-- {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		QuickSort(nums)
	}
}

func BenchmarkQuickSortInOrder(b *testing.B) {
	nums := []int{}
	for i := -1000; i <= 1000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		QuickSort(nums)
	}
}
