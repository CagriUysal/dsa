// Properties:
//  1. Complete Binary Tree: All levels are filled except possibly the last level,
//     which is filled from left to right.
//  2. Heap Property: For any given node i,
//     - Value of node i is less than or equal to its children
//     - A[parent(i)] <= A[i]
//     - Root contains the minimum element
package heap

import (
	"math"
)

type MinHeap struct {
	Heap
}

func (h *MinHeap) BuildMinHeap(arr []int) {
	h.Arr = arr
	h.HeapSize = len(arr)

	lastNonLeafNode := (h.HeapSize / 2) - 1
	for i := lastNonLeafNode; i >= 0; i-- {
		h.heapify(i)
	}
}

func (h *MinHeap) heapify(i int) {
	l := h.left(i)
	r := h.right(i)

	minimum := i
	if l < h.HeapSize && h.Arr[l] < h.Arr[minimum] {
		minimum = l
	}

	if r < h.HeapSize && h.Arr[r] < h.Arr[minimum] {
		minimum = r
	}

	if minimum != i {
		h.swap(i, minimum)
		h.heapify(minimum)
	}
}

func (h *MinHeap) GetMin() int {
	if h.HeapSize < 1 {
		panic("heap underflow")
	}

	return h.Arr[0]
}

func (h *MinHeap) ExtractMin() int {
	min := h.GetMin()

	h.swap(0, h.HeapSize-1)
	h.HeapSize--

	h.heapify(0)

	return min
}

func (h *MinHeap) DecreaseKey(i, k int) {
	if k > h.Arr[i] {
		panic("new key is greater than current key")
	}

	h.Arr[i] = k

	for i > 0 {
		p := h.parent(i)
		if h.Arr[p] < h.Arr[i] {
			break
		}
		h.swap(i, p)
		i = p
	}
}

func (h *MinHeap) Insert(k int) {
	if h.HeapSize == len(h.Arr) {
		h.Arr = append(h.Arr, math.MaxInt)
	} else {
		h.Arr[h.HeapSize] = math.MaxInt
	}

	h.HeapSize++
	h.DecreaseKey(h.HeapSize-1, k)
}
