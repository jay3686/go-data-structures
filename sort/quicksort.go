// Package sort implements the quick sort algorithm.

package sort

// QuickSort sorts an array of ints in ascending order.
func QuickSort(a []int) {
	quickSort(a, 0, len(a))
}

// quickSort recursivelly sorts indexes s-e of an array
func quickSort(a []int, s int, e int) {
	// choose middle element as pivot.
	p := a[(s+e)/2]

	if s == e-1 {
		return
	}
	smallest, i, j := 0, 0, len(a)-1

	for i < j {
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
	quickSort(a, s, (s+e)/2)
	quickSort(a, (s+e)/2, e)

}
