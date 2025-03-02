package heap_test

import (
	"dsa/ds/heap"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildMaxHeap(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := heap.MaxHeap{}

	maxHeap.BuildMaxHeap(input)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMaxHeapHeapify(t *testing.T) {
	input := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	maxHeap := heap.MaxHeap{
		heap.Heap{
			Arr:      input,
			HeapSize: len(input),
		},
	}

	maxHeap.Heapify(1)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMapHeapGetMax(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := heap.MaxHeap{}
	maxHeap.BuildMaxHeap(input)

	expected := 16
	result := maxHeap.GetMax()

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestMaxHeapExtractMax(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := heap.MaxHeap{}
	maxHeap.BuildMaxHeap(input)

	expected := 16
	result := maxHeap.ExtractMax()

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(9, maxHeap.HeapSize); diff != "" {
		t.Error(diff)
	}
}

func TestMaxHeapIncreaseKey(t *testing.T) {
	maxHeap := heap.MaxHeap{
		heap.Heap{
			Arr:      []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			HeapSize: 10,
		},
	}

	maxHeap.IncreaseKey(8, 15)

	expected := []int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}
	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMaxHeapInsert(t *testing.T) {
	maxHeap := heap.MaxHeap{
		heap.Heap{
			Arr:      []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
			HeapSize: 10,
		},
	}

	maxHeap.Insert(17)

	expected := []int{17, 16, 10, 8, 14, 9, 3, 2, 4, 1, 7}
	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestBuildMinHeap(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	minHeap := heap.MinHeap{}

	minHeap.BuildMinHeap(input)

	expected := []int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16}

	if diff := cmp.Diff(expected, minHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMinHeapGetMin(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	minHeap := heap.MinHeap{}
	minHeap.BuildMinHeap(input)

	expected := 1
	result := minHeap.GetMin()

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestMinHeapExtractMin(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	minHeap := heap.MinHeap{}
	minHeap.BuildMinHeap(input)

	expected := 1
	result := minHeap.ExtractMin()

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

	if diff := cmp.Diff(9, minHeap.HeapSize); diff != "" {
		t.Error(diff)
	}
}

func TestMinHeapDecreaseKey(t *testing.T) {
	minHeap := heap.MinHeap{
		heap.Heap{
			Arr:      []int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			HeapSize: 10,
		},
	}

	minHeap.DecreaseKey(8, 0)

	expected := []int{0, 1, 3, 2, 7, 9, 10, 14, 4, 16}
	if diff := cmp.Diff(expected, minHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMinHeapInsert(t *testing.T) {
	minHeap := heap.MinHeap{
		heap.Heap{
			Arr:      []int{1, 2, 3, 4, 7, 9, 10, 14, 8, 16},
			HeapSize: 10,
		},
	}

	minHeap.Insert(0)

	expected := []int{0, 1, 3, 4, 2, 9, 10, 14, 8, 16, 7}
	if diff := cmp.Diff(expected, minHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

