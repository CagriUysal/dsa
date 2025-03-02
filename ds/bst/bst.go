package bst

/*
Binary Search Tree (BST) Properties:

1. Each node contains a unique key (no duplicates)
2. For any node:
  - All keys in the left subtree are less than the node's key
  - All keys in the right subtree are greater than the node's key
3. The left and right subtrees are also binary search trees
4. Operations:
  - Search: O(h) time, where h is height of tree
  - Insert: O(h) time
  - Delete: O(h) time
  - Best case (balanced): h = log n
  - Worst case (unbalanced): h = n
*/

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type Bst struct {
	Root *Node
}

func (t *Bst) InOrderWalk(node *Node, keys *[]int) {
	if node == nil {
		return
	}

	t.InOrderWalk(node.Left, keys)
	*keys = append(*keys, node.Key)
	t.InOrderWalk(node.Right, keys)
}

func (t *Bst) PreOrderWalk(node *Node, keys *[]int) {
	if node == nil {
		return
	}

	*keys = append(*keys, node.Key)
	t.PreOrderWalk(node.Left, keys)
	t.PreOrderWalk(node.Right, keys)
}

func (t *Bst) PostOrderWalk(node *Node, keys *[]int) {
	if node == nil {
		return
	}

	t.PostOrderWalk(node.Left, keys)
	t.PostOrderWalk(node.Right, keys)
	*keys = append(*keys, node.Key)
}

func (t *Bst) Search(node *Node, key int) *Node {

	if node == nil || node.Key == key {
		return node
	}

	if key > node.Key {
		return t.Search(node.Right, key)
	} else {
		return t.Search(node.Left, key)
	}
}

func (t *Bst) SearchIteratively(key int) *Node {
	curr := t.Root

	for curr != nil && curr.Key != key {
		if key > curr.Key {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	return curr
}

func (t *Bst) Minimum(node *Node) *Node {
	curr := node

	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func (t *Bst) Maximum(node *Node) *Node {
	curr := node

	for curr.Right != nil {
		curr = curr.Right
	}

	return curr
}

func (t *Bst) Successor(node *Node) *Node {
	if node.Right != nil {
		return t.Minimum(node.Right)
	}

	p := node.Parent
	for p != nil && p.Right == node {
		node = p
		p = p.Parent
	}

	return p
}

func (t *Bst) Predecessor(node *Node) *Node {
	if node.Left != nil {
		return t.Maximum(node.Left)
	}

	p := node.Parent
	for p != nil && p.Left == node {
		node = p
		p = p.Parent
	}

	return p
}

func (t *Bst) Insert(node *Node) {
	curr := t.Root
	var trailing *Node
	for curr != nil {
		trailing = curr
		if node.Key > curr.Key {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	if trailing == nil {
		t.Root = node
		return
	}

	node.Parent = trailing
	if node.Key > trailing.Key {
		trailing.Right = node
	} else {
		trailing.Left = node
	}
}

func (t *Bst) Delete(node *Node) {
	if node.Left == nil {
		t.transplant(node, node.Right)
		return
	} else if node.Right == nil {
		t.transplant(node, node.Left)
		return
	}

	successor := t.Successor(node)
	if successor != node.Right {
		t.transplant(successor, successor.Right)
		successor.Right = node.Right
		successor.Right.Parent = successor
	}

	t.transplant(node, successor)
	successor.Left = node.Left
	successor.Left.Parent = successor
}

// replace `destination` node with `target` node
// it does not attemps to update `target.left` or `target.right`
func (t *Bst) transplant(destination *Node, target *Node) {
	if destination.Parent == nil {
		t.Root = target
	} else if destination.Parent.Left == destination {
		destination.Parent.Left = target
	} else {
		destination.Parent.Right = target
	}

	if target != nil {
		target.Parent = destination.Parent
	}
}
