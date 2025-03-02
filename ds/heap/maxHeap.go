// Properties:
//  1. Complete Binary Tree: All levels are filled except possibly the last level,
//     which is filled from left to right.
//  2. Heap Property: For any given node i,
//     - Value of node i is greater than or equal to its children
//     - A[parent(i)] ≥ A[i]
//     - Root contains the maximum element

package heap

import (
	"math"
)

type MaxHeap struct {
	Heap
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
	l := h.left(i)
	r := h.right(i)

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
		p := h.parent(i)
		if h.Arr[p] > h.Arr[i] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MaxHeap) Insert(k int) {
	if h.HeapSize == len(h.Arr) {
		h.Arr = append(h.Arr, math.MinInt)
	} else {
		h.Arr[h.HeapSize] = math.MinInt
	}

	h.HeapSize++
	h.IncreaseKey(h.HeapSize-1, k)
}
