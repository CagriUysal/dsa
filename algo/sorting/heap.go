package sorting

import (
	heap "dsa/ds/heap"
)

func HeapSort(arr []int) []int {
	maxHeap := heap.MaxHeap{}
	maxHeap.BuildMaxHeap(arr)

	for i := len(maxHeap.Arr) - 1; i >= 0; i -= 1 {
		maxHeap.Arr[i], maxHeap.Arr[0] = maxHeap.Arr[0], maxHeap.Arr[i]
		maxHeap.HeapSize = i
		maxHeap.Heapify(0)
	}

	return maxHeap.Arr
}
