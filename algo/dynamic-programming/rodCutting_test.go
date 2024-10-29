package dynamicprogramming_test

import (
	dynamicprogramming "dsa/algo/dynamic-programming"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRodCutNaive(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 10

	expected := 30
	result := dynamicprogramming.RodCutNaive(prices, n)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestMemoizedRodCut(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 10

	expected := 30
	result := dynamicprogramming.MemoizedRodCut(prices, n)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestBottomUpRobCut(t *testing.T) {
	prices := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	n := 10

	expected := 30
	result := dynamicprogramming.BottomUpRobCut(prices, n)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
