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

func BenchmarkBubbleSortReverseOrder(b *testing.B) {
	nums := []int{}
	for i := 1000; i >= -1000; i-- {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkBubbleSort2ReverseOrder(b *testing.B) {
	nums := []int{}
	for i := 1000; i >= -1000; i-- {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort2(nums)
	}
}

func BenchmarkBubbleSortInOrder(b *testing.B) {
	nums := []int{}
	for i := -1000; i <= 1000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkBubbleSort2InOrder(b *testing.B) {
	nums := []int{}
	for i := -1000; i <= 1000; i++ {
		nums = append(nums, i)
	}
	for i := 0; i < b.N; i++ {
		BubbleSort2(nums)
	}
}
