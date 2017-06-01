// Package heap implements heap based data structures.
package heap

// MinHeap is a class that implements a Min Heap data structure
type MinHeap struct {
	data []int
}

// Len returns the length of the heap
func (h MinHeap) Len() int { return len(h.data) }

// FindMin returns the minimum element in the heap
func (h MinHeap) FindMin() int {
	if len(h.data) == 0 {
		//TODO:  Figure out how to cleanly handle this case
		return 0
	}
	return h.data[0]
}

// ExtractMin deletes and returns the min element from the heap
func (h *MinHeap) ExtractMin() int {
	if len(h.data) == 0 {
		//TODO:  Figure out how to cleanly handle this case
		return 0
	}
	n := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]

	h.heapify(0)
	return n
}

// heapify assures that the node at i and its children satisfy the min heap property
func (h MinHeap) heapify(i int) {

	// children of index i are at 2i + 1 and 2i + 2
	leftChild := 2*i + 1
	rightChild := 2*i + 2

	if leftChild < len(h.data) && h.data[leftChild] < h.data[i] {
		h.data[leftChild], h.data[i] = h.data[i], h.data[leftChild]
		h.heapify(leftChild)
	}
	if rightChild < len(h.data) && h.data[rightChild] < h.data[i] {
		h.data[rightChild], h.data[i] = h.data[i], h.data[rightChild]
		h.heapify(rightChild)
	}
}

// Insert adds an element into the heap such that the heap property is still satsified
func (h *MinHeap) Insert(n int) {

	h.data = append(h.data[:len(h.data)], n)
	// swap with parent if child is larger until at root node
	// parent of index i is  i / 2
	c := len(h.data) - 1
	p := (c - 1) / 2
	for c > 0 && h.data[p] > h.data[c] {
		h.data[p], h.data[c] = h.data[c], h.data[p]
		c, p = p, (p-1)/2
	}
}
