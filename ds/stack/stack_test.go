package stack_test

import (
	"dsa/ds/stack"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestIsEmpty(t *testing.T) {
	stack := stack.Stack{}

	if diff := cmp.Diff(true, stack.IsEmpty()); diff != "" {
		t.Error(diff)
	}

	stack.Push(1)
	stack.Push(2)

	if diff := cmp.Diff(false, stack.IsEmpty()); diff != "" {
		t.Error(diff)
	}
}

func TestPopPush(t *testing.T) {
	stack := stack.Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, _ := stack.Pop()

	if diff := cmp.Diff(item, 3); diff != "" {
		t.Error(diff)
	}
}
