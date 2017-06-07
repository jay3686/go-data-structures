// Package Merge implements the merge sort algorithm.

package sort

// MergeSort sorts an array of ints in ascending order.
func MergeSort(a []int) {
	mergeSort(a, 0, len(a))
}

// mergeSort recursivelly sorts indexes s-e of an array
func mergeSort(a []int, s int, e int) []int {
	mid := (s + e) / 2
	if s == e-1 {
		return []int{a[s]}
	}
	// sort left and right
	left := mergeSort(a, s, mid)
	right := mergeSort(a, mid, e)
	// merge
	li, ri := 0, 0
	for i := s; i < e; i++ {
		if li >= len(left) {
			for ri < len(right) {
				a[i] = right[ri]
				i++
				ri++
			}
		} else if ri >= len(right) {
			for li < len(left) {
				a[i] = left[li]
				i++
				li++
			}
		} else if right[ri] < left[li] {
			a[i] = right[ri]
			ri++
		} else {
			a[i] = left[li]
			li++
		}
	}

	r := make([]int, e-s)
	copy(r, a[s:e])
	return r
}
