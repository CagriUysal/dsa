package maxHeap

type MaxHeap struct {
	Arr      []int
	HeapSize int
}

func (h *MaxHeap) BuildMaxHeap(arr []int) {
	h.Arr = arr
	h.HeapSize = len(h.Arr)
	for i := (h.HeapSize / 2) - 1; i >= 0; i -= 1 {
		h.MaxHeapify(i)
	}
}

// heapify ith index
// assumes left(i) and right(i) are valid heaps
func (h *MaxHeap) MaxHeapify(i int) {
	l := left(i)
	r := right(i)

	largest := i
	if l < h.HeapSize && h.Arr[l] > h.Arr[largest] {
		largest = l
	}

	if r < h.HeapSize && h.Arr[r] > h.Arr[largest] {
		largest = r
	}

	if largest != i {
		h.Arr[largest], h.Arr[i] = h.Arr[i], h.Arr[largest]
		h.MaxHeapify(largest)
	}
}

// give the parent index in the heap
func parent(i int) int {
	return i / 2
}

// give the left index in the heap
func left(i int) int {
	return 2*i + 1
}

// give the right index in the heap
func right(i int) int {
	return 2*i + 2
}
