package bst_test

import (
	"dsa/ds/bst"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var root *bst.Node = &bst.Node{
	Key: 6,
	Left: &bst.Node{
		Key: 5,
		Left: &bst.Node{
			Key: 2,
		},
		Right: &bst.Node{
			Key: 5,
		},
	},
	Right: &bst.Node{
		Key: 7,
		Right: &bst.Node{
			Key: 8,
		},
	},
}

func newNode(key int) *bst.Node {
	return &bst.Node{Key: key}
}

func TestInOrderWalk(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	var keys []int
	tree.InOrderWalk(tree.Root, &keys)
	expected := []int{2, 5, 5, 6, 7, 8}
	if diff := cmp.Diff(keys, expected); diff != "" {
		t.Error(diff)
	}
}

func TestPreOrderWalk(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	var keys []int
	tree.PreOrderWalk(tree.Root, &keys)
	expected := []int{6, 5, 2, 5, 7, 8}
	if diff := cmp.Diff(keys, expected); diff != "" {
		t.Error(diff)
	}
}

func TestPostOrderWalk(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	var keys []int
	tree.PostOrderWalk(tree.Root, &keys)
	expected := []int{2, 5, 5, 8, 7, 6}
	if diff := cmp.Diff(keys, expected); diff != "" {
		t.Error(diff)
	}
}

func TestSearch_Found(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	key := 8
	result := tree.Search(tree.Root, key)
	if diff := cmp.Diff(key, result.Key); diff != "" {
		t.Error(diff)
	}
}

func TestSearch_NotFound(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	key := 9
	result := tree.Search(tree.Root, key)
	if diff := cmp.Diff((*bst.Node)(nil), result); diff != "" {
		t.Error(diff)
	}
}

func TestSearchIteratively_Found(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	key := 5
	result := tree.SearchIteratively(key)
	if diff := cmp.Diff(key, result.Key); diff != "" {
		t.Error(diff)
	}
}

func TestSearchIteratively_NotFound(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	key := 11
	result := tree.SearchIteratively(key)
	if diff := cmp.Diff((*bst.Node)(nil), result); diff != "" {
		t.Error(diff)
	}
}

func TestMiniumum(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	result := tree.Minimum(tree.Root)
	if diff := cmp.Diff(2, result.Key); diff != "" {
		t.Error(diff)
	}
}

func TestMaximum(t *testing.T) {
	tree := bst.Bst{
		Root: root,
	}

	result := tree.Maximum(tree.Root)
	if diff := cmp.Diff(8, result.Key); diff != "" {
		t.Error(diff)
	}
}

func TestSuccessor(t *testing.T) {
	// Manually constructing the BST
	//       5
	//     /   \
	//    3     8
	//   / \   / \
	//  2   4 6   9
	root := newNode(5)
	root.Left = newNode(3)
	root.Right = newNode(8)
	root.Left.Parent = root
	root.Right.Parent = root
	root.Left.Left = newNode(2)
	root.Left.Right = newNode(4)
	root.Left.Left.Parent = root.Left
	root.Left.Right.Parent = root.Left
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(9)
	root.Right.Left.Parent = root.Right
	root.Right.Right.Parent = root.Right

	tree := bst.Bst{Root: root}

	testCases := []struct {
		node        *bst.Node
		expectedKey int
		expectedNil bool
	}{
		{node: root.Left.Right, expectedKey: 5, expectedNil: false}, // Successor of 4 is 5
		{node: root.Right, expectedKey: 9, expectedNil: false},      // Successor of 8 is 9
		{node: root.Right.Right, expectedKey: 0, expectedNil: true}, // 9 has no successor
	}

	for _, tc := range testCases {
		actual := tree.Successor(tc.node)
		if tc.expectedNil && actual != nil {
			t.Errorf("Expected nil, got %v for node with key %d", actual, tc.node.Key)
		} else if !tc.expectedNil {
			if actual == nil {
				t.Errorf("Expected successor with key %d, got nil", tc.expectedKey)
			} else if diff := cmp.Diff(tc.expectedKey, actual.Key); diff != "" {
				t.Error(diff)
			}
		}
	}
}

func TestPredecessor(t *testing.T) {
	// Manually constructing the BST
	//       5
	//     /   \
	//    3     8
	//   / \   / \
	//  2   4 6   9
	root := newNode(5)
	root.Left = newNode(3)
	root.Right = newNode(8)
	root.Left.Parent = root
	root.Right.Parent = root
	root.Left.Left = newNode(2)
	root.Left.Right = newNode(4)
	root.Left.Left.Parent = root.Left
	root.Left.Right.Parent = root.Left
	root.Right.Left = newNode(6)
	root.Right.Right = newNode(9)
	root.Right.Left.Parent = root.Right
	root.Right.Right.Parent = root.Right

	tree := bst.Bst{Root: root}

	// Testing Predecessor function
	testCases := []struct {
		node        *bst.Node
		expectedKey int
		expectedNil bool
	}{
		{node: root.Right.Right, expectedKey: 8, expectedNil: false}, // Predecessor of 9 is 8
		{node: root, expectedKey: 4, expectedNil: false},             // Predecessor of 5 is 4
		{node: root.Left.Left, expectedKey: 0, expectedNil: true},    // 2 has no predecessor
	}

	for _, tc := range testCases {
		actual := tree.Predecessor(tc.node)
		if tc.expectedNil && actual != nil {
			t.Errorf("Expected nil, got %v for node with key %d", actual, tc.node.Key)
		} else if !tc.expectedNil {
			if actual == nil {
				t.Errorf("Expected predecessor with key %d, got nil", tc.expectedKey)
			} else if diff := cmp.Diff(tc.expectedKey, actual.Key); diff != "" {
				t.Errorf("Predecessor key mismatch (-want +got):\n%s", diff)
			}
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name           string
		keys           []int
		expectedRoot   int
		expectedParent map[int]int
	}{
		{
			name:           "Insert into empty tree",
			keys:           []int{10},
			expectedRoot:   10,
			expectedParent: map[int]int{},
		},
		{
			name:           "Insert multiple ascending",
			keys:           []int{10, 20, 30},
			expectedRoot:   10,
			expectedParent: map[int]int{20: 10, 30: 20},
		},
		{
			name:           "Insert multiple descending",
			keys:           []int{30, 20, 10},
			expectedRoot:   30,
			expectedParent: map[int]int{20: 30, 10: 20},
		},
		{
			name:           "Insert mixed",
			keys:           []int{20, 10, 30},
			expectedRoot:   20,
			expectedParent: map[int]int{10: 20, 30: 20},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tree := &bst.Bst{}
			for _, key := range tc.keys {
				tree.Insert(newNode(key))
			}

			if diff := cmp.Diff(tc.expectedRoot, tree.Root.Key); diff != "" {
				t.Error(diff)
			}

			for childKey, expectedParentKey := range tc.expectedParent {
				node := tree.SearchIteratively(childKey)
				if node == nil || node.Parent == nil {
					t.Fatalf("node %d not found or has no parent", childKey)
				}
				if diff := cmp.Diff(expectedParentKey, node.Parent.Key); diff != "" {
					t.Error(diff)
				}
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name          string
		insertKeys    []int
		deleteKey     int
		expectedOrder []int
	}{
		{
			name:          "Delete leaf node",
			insertKeys:    []int{10, 5, 15},
			deleteKey:     5,
			expectedOrder: []int{10, 15},
		},
		{
			name:          "Delete node with one child",
			insertKeys:    []int{10, 5, 15, 12},
			deleteKey:     15,
			expectedOrder: []int{5, 10, 12},
		},
		{
			name:          "Delete node with two children",
			insertKeys:    []int{10, 5, 15, 12, 17},
			deleteKey:     15,
			expectedOrder: []int{5, 10, 12, 17},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tree := createBSTFromSlice(tc.insertKeys)
			nodeToDelete := tree.SearchIteratively(tc.deleteKey)
			if nodeToDelete == nil {
				t.Fatalf("Node %d not found for deletion", tc.deleteKey)
			}

			tree.Delete(nodeToDelete)

			var gotOrder []int
			tree.InOrderWalk(tree.Root, &gotOrder)
			if diff := cmp.Diff(tc.expectedOrder, gotOrder); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func createBSTFromSlice(keys []int) *bst.Bst {
	tree := &bst.Bst{}
	for _, key := range keys {
		tree.Insert(newNode(key))
	}
	return tree
}
