package redblacktree_test

import (
	rb "dsa/ds/red-black-tree"
	"testing"
)

func newRedBlackTree() *rb.RedBlackTree {
	NIL := &rb.Node{}
	return &rb.RedBlackTree{
		Root: NIL,
		NIL:  NIL,
	}
}

func newNode(value int) *rb.Node {
	return &rb.Node{
		Value: value,
	}
}

func isRed(node *rb.Node) bool {
	if node == nil {
		return false
	}
	return node.Color == rb.RED
}

func isBlack(node *rb.Node) bool {
	if node == nil {
		return true
	}
	return node.Color == rb.BLACK
}

func validateRBTree(t *testing.T, node *rb.Node, blackCount int, pathBlackCount *int) bool {
	if node == nil {
		if *pathBlackCount == -1 {
			*pathBlackCount = blackCount
		}
		return *pathBlackCount == blackCount
	}

	if isRed(node) {
		if isRed(node.Left) || isRed(node.Right) {
			t.Errorf("Red violation at node %d", node.Value)
			return false
		}
	}

	if isBlack(node) {
		blackCount++
	}

	return validateRBTree(t, node.Left, blackCount, pathBlackCount) &&
		validateRBTree(t, node.Right, blackCount, pathBlackCount)
}

func TestInsert(t *testing.T) {
	tree := newRedBlackTree()
	values := []int{10, 20, 30, 15, 25, 5, 1}

	for _, v := range values {
		tree.Insert(newNode(v))
	}

	pathBlackCount := -1
	if !validateRBTree(t, tree.Root, 0, &pathBlackCount) {
		t.Errorf("Red-Black Tree property validation failed")
	}

	extraValues := []int{12, 17, 22, 27, 3, 8}
	for _, v := range extraValues {
		tree.Insert(newNode(v))
		pathBlackCount = -1
		if !validateRBTree(t, tree.Root, 0, &pathBlackCount) {
			t.Errorf("Red-Black Tree property validation failed after inserting %d", v)
		}
	}
}
