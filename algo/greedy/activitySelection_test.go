package greedy_test

import (
	greedy "dsa/algo/greedy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRecursiveActivitySelection(t *testing.T) {
	startTimes := []int{1, 3, 0, 5, 3, 5, 6, 7, 8, 2, 12}
	finishTimes := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}

	expected := []int{0, 3, 7, 10}
	result := greedy.RecursiveActivitySelection(startTimes, finishTimes)


	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}

func TestIterativeActivitySelection(t *testing.T) {
	startTimes := []int{1, 3, 0, 5, 3, 5, 6, 7, 8, 2, 12}
	finishTimes := []int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}

	expected := []int{0, 3, 7, 10}
	result := greedy.ActivitySelection(startTimes, finishTimes)

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Error(diff)
	}
}
