// Package sort implements the quick sort algorithm.

package sort

// QuickSort sorts an array of ints in ascending order.
func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

// quickSort recursivelly sorts indexes s-e of an array
func quickSort(a []int, s int, e int) {
	// choose middle element as pivot.
	p := a[(s+e)/2]
	if s >= e {
		return
	}

	smallest, i, j := s, s, e

	for i <= j {
		if a[i] < p {
			a[smallest], a[i] = a[i], a[smallest]
			i++
			smallest++
		} else if a[i] > p {
			a[j], a[i] = a[i], a[j]
			j--
		} else if a[i] == p {
			i++
		}
	}
	left := smallest - 1
	// if smallest < 0 {
	// 	left = 0
	// }
	right := j + 1
	// if j >= len(a) {
	// 	j = len(a) - 1
	// }
	quickSort(a, s, left)
	quickSort(a, right, e)

}
