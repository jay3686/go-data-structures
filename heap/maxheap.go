// Package heap implements heap based data structures.
package heap

// MaxHeap is a class that implements a Max Heap data structure
type MaxHeap struct {
	data []int
}

// Len returns the length of the heap
func Len(h MaxHeap) int { return len(h.data) }

// FindMax returns the maximum element in the heap
func (h MaxHeap) FindMax() int {
	if len(h.data) == 0 {
		//TODO:  Figure out how to cleanly handle this case
		return 0
	}
	return h.data[0]
}

// ExtractMax deletes and returns the max element from the heap
func (h MaxHeap) ExtractMax() int {
	if len(h.data) == 0 {
		//TODO:  Figure out how to cleanly handle this case
		return 0
	}
	n := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	// TODO, rebalance by swapping eleement down

	return n
}

// heapify assures that the node at i and its children satisfy the max heap property
func (h MaxHeap) heapify(i int) {

	// children of index i are at 2i + 1 and 2i + 2
	leftChild := 2*i + 1
	rightChild := 2*i + 2

	if leftChild < len(h.data) && h.data[leftChild] > h.data[i] {
		h.data[leftChild], h.data[i] = h.data[i], h.data[leftChild]
		h.heapify(leftChild)
	}
	if rightChild < len(h.data) && h.data[rightChild] > h.data[i] {
		h.data[rightChild], h.data[i] = h.data[i], h.data[rightChild]
		h.heapify(rightChild)
	}
}

// Insert adds an element into the heap such that the heap property is still stasified
func (h MaxHeap) Insert(n int) {
	h.data = append(h.data[:len(h.data)-1], n)

	// swap with parent if child is larger until at root node
	// parent of index i is  i / 2
	c := len(h.data) - 1
	p := c / 2
	for c > 0 && h.data[p] < h.data[c] {
		h.data[p], h.data[c] = h.data[c], h.data[p]
		p, c = c, c/2
	}
}
