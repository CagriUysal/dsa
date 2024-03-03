package sorting_test

import (
	"dsa/src/sorting"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var input []int
var expected = []int{5, 12, 23, 34, 45, 56, 67, 78, 89, 90}

func setup() {
	input = []int{23, 45, 67, 12, 89, 34, 56, 78, 90, 5}
}

func TestBubbleSort(t *testing.T) {
	setup()
	result := sorting.BubbleSort(input)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

}

func TestInsertionSort(t *testing.T) {
	setup()
	result := sorting.InsertionSort(input)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}

}

func TestMergeSort(t *testing.T) {
	setup()

	result := sorting.MergeSort(input)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestHeapSort(t *testing.T) {
	setup()

	result := sorting.HeapSort(input)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestQuickSort(t *testing.T) {
	setup()

	result := sorting.QuickSort(input)
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestCountSort(t *testing.T) {
	input := []int{2, 5, 3, 0, 2, 3, 0, 3}

	result := sorting.CountingSort(input, 5)
	expected := []int{0, 0, 2, 2, 3, 3, 3, 5}
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestRadixSort(t *testing.T) {
	input := []int{329, 547, 657, 839, 436, 720, 355, 7}

	result := sorting.RadixSort(input, 3)
	expected := []int{7, 329, 355, 436, 547, 657, 720, 839}
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
