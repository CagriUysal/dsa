package maxHeap

import (
	"math"
)

type MaxHeap struct {
	Arr      []int
	HeapSize int
}

func (h *MaxHeap) BuildMaxHeap(arr []int) {
	h.Arr = arr
	h.HeapSize = len(h.Arr)
	for i := (h.HeapSize / 2) - 1; i >= 0; i -= 1 {
		h.Heapify(i)
	}
}

// heapify ith index
// assumes left(i) and right(i) are valid heaps
func (h *MaxHeap) Heapify(i int) {
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
		h.swap(largest, i)
		h.Heapify(largest)
	}
}

func (h *MaxHeap) GetMax() int {
	if h.HeapSize < 1 {
		panic("heap underflow")
	}

	return h.Arr[0]
}

func (h *MaxHeap) ExtractMax() int {
	max := h.GetMax()

	h.swap(0, h.HeapSize-1)
	h.HeapSize -= 1

	h.Heapify(0)

	return max
}

func (h *MaxHeap) IncreaseKey(i, k int) {
	if h.Arr[i] > k {
		panic("current key is greater!")
	}

	h.Arr[i] = k

	for i > 0 {
		p := parent(i)
		if h.Arr[p] > h.Arr[i] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MaxHeap) Insert(k int) {
	h.HeapSize += 1

	minInt := math.MinInt
	h.Arr = append(h.Arr, minInt)

	h.IncreaseKey(h.HeapSize-1, k)
}

func (h *MaxHeap) swap(i, j int) {
	h.Arr[i], h.Arr[j] = h.Arr[j], h.Arr[i]
}

// give the parent index in the heap
func parent(i int) int {
	return (i - 1) / 2
}

// give the left index in the heap
func left(i int) int {
	return 2*i + 1
}

// give the right index in the heap
func right(i int) int {
	return 2*i + 2
}
