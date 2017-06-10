package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	nums := []int{6, 5, 4, 3, 2, 1}
	MergeSort(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 3, 4, 5, 6}) {
		t.Error("Not sorted in asc order", nums)
	}
}

func BenchmarkMergeSortReverseOrder(b *testing.B) {
	nums := []int{}
	for i := 1000; i >= -1000; i-- {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		MergeSort(nums)
	}
}

func BenchmarkMergeSortInOrder(b *testing.B) {
	nums := []int{}
	for i := -1000; i <= 1000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		MergeSort(nums)
	}
}
