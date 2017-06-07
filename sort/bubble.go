// Package bubble implements the bubble sort algorithm.

package sort

// BubbleSort sorts an array of ints in ascending order.
func BubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// BubbleSort2 sorts an array of ints in ascending order.
func BubbleSort2(a []int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
