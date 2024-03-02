package maxHeap_test

import (
	"dsa/src/maxHeap"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBuildMaxHeap(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := maxHeap.MaxHeap{}

	maxHeap.BuildMaxHeap(input)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestMaxHeapify(t *testing.T) {
	input := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	maxHeap := maxHeap.MaxHeap{
		Arr:      input,
		HeapSize: len(input),
	}

	maxHeap.MaxHeapify(1)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}
