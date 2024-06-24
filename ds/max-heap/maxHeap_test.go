package maxHeap_test

import (
	"dsa/ds/max-heap"
	"fmt"
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

func TestHeapify(t *testing.T) {
	input := []int{16, 4, 10, 14, 7, 9, 3, 2, 8, 1}
	maxHeap := maxHeap.MaxHeap{
		Arr:      input,
		HeapSize: len(input),
	}

	maxHeap.Heapify(1)

	expected := []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1}

	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestGetMax(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := maxHeap.MaxHeap{}
	maxHeap.BuildMaxHeap(input)

	expected := 16
	result := maxHeap.GetMax()

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestExtractMax(t *testing.T) {
	input := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	maxHeap := maxHeap.MaxHeap{}
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

func TestIncreaseKey(t *testing.T) {
	maxHeap := maxHeap.MaxHeap{
		Arr:      []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		HeapSize: 10,
	}

	maxHeap.IncreaseKey(8, 15)

	expected := []int{16, 15, 10, 14, 7, 9, 3, 2, 8, 1}
	fmt.Println(maxHeap.Arr)
	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}

func TestInsert(t *testing.T) {
	maxHeap := maxHeap.MaxHeap{
		Arr:      []int{16, 14, 10, 8, 7, 9, 3, 2, 4, 1},
		HeapSize: 10,
	}

	maxHeap.Insert(17)

	expected := []int{17, 16, 10, 8, 14, 9, 3, 2, 4, 1, 7}
	if diff := cmp.Diff(expected, maxHeap.Arr); diff != "" {
		t.Error(diff)
	}
}
