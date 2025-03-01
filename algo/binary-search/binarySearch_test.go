package binarysearch_test

import (
	"dsa/algo/binary-search"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBinarySearchFound(t *testing.T) {
	nums := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	key := 20

	result := binarysearch.BinarySearch(nums, key)

	expected := 7
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestBinarySearchNotFound(t *testing.T) {
	nums := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	key := 42

	result := binarysearch.BinarySearch(nums, key)

	expected := -1
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
