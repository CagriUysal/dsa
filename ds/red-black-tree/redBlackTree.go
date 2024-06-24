// properties of red black trees
// 1. Every node is either red or black.
// 2. The root is black.
// 3. Every leaf (NIL) is black.
// 4. If a node is red, then both its children are black.
// 5. For each node, all simple paths from the node to descendant leaves contain the same number of black nodes.

package redblacktree

import "fmt"

const (
	RED   = true
	BLACK = false
)

type Node struct {
	Color bool
	Value int
	Left  *Node
	Right *Node
	P     *Node
}

type RedBlackTree struct {
	Root *Node
	NIL  *Node
}

func (tree *RedBlackTree) leftRotate(x *Node) {
	y := x.Right

	x.Right = y.Left
	if y.Left != tree.NIL {
		y.Left.P = x
	}

	y.P = x.P
	if x.P == tree.NIL {
		tree.Root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}

	x.P = y
	y.Left = x
}

func (tree *RedBlackTree) rightRotate(x *Node) {
	y := x.Left

	x.Left = y.Right
	if y.Right != tree.NIL {
		y.Right.P = x
	}

	y.P = x.P
	if x.P == tree.NIL {
		tree.Root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}

	x.P = y
	y.Right = x
}

func (tree *RedBlackTree) Insert(node *Node) {
	x := tree.Root
	y := tree.NIL

	for x != tree.NIL {
		y = x
		if node.Value > x.Value {
			x = x.Right
		} else {
			x = x.Left
		}
	}

	node.P = y
	if y == tree.NIL {
		tree.Root = node
	} else if node.Value > y.Value {
		y.Right = node
	} else {
		y.Left = node
	}

	node.Left = tree.NIL
	node.Right = tree.NIL
	node.Color = RED

	tree.insertFixup(node)
}

func (tree *RedBlackTree) insertFixup(node *Node) {
	for node.P.Color == RED {
		if node.P == node.P.P.Left {
			uncle := node.P.P.Right

			if uncle.Color == RED {
				node.P.Color = BLACK
				uncle.Color = BLACK
				node.P.P.Color = RED
				node = node.P.P
			} else {
				if node == node.P.Right {
					node = node.P
					tree.leftRotate(node)
				}
				node.P.Color = BLACK
				node.P.P.Color = RED
				tree.rightRotate(node.P.P)
			}
		} else {
			uncle := node.P.P.Left

			if uncle.Color == RED {
				node.P.Color = BLACK
				uncle.Color = BLACK
				node.P.P.Color = RED
				node = node.P.P
			} else {
				if node == node.P.Left {
					node = node.P
					tree.rightRotate(node)
				} else {
					node.P.Color = BLACK
					node.P.P.Color = RED
					tree.leftRotate(node.P.P)
				}
			}
		}
	}

	tree.Root.Color = BLACK
}

func (tree *RedBlackTree) PrettyPrint() {
	if tree.Root == tree.NIL {
		fmt.Println("Empty tree")
		return
	}

	type NodeLevel struct {
		node  *Node
		level int
	}
	queue := []NodeLevel{{node: tree.Root, level: 0}}
	currentLevel := 0
	for len(queue) > 0 {
		nl := queue[0]
		queue = queue[1:]

		if nl.level > currentLevel {
			currentLevel = nl.level
			fmt.Println()
		}

		nodeColor := "B"
		if nl.node.Color == RED {
			nodeColor = "R"
		}
		fmt.Printf("%d%s ", nl.node.Value, nodeColor)

		if nl.node.Left != tree.NIL {
			queue = append(queue, NodeLevel{nl.node.Left, nl.level + 1})
		}
		if nl.node.Right != tree.NIL {
			queue = append(queue, NodeLevel{nl.node.Right, nl.level + 1})
		}
	}
	fmt.Println()
}

func (tree *RedBlackTree) Delete(z *Node) {
	y := z
	yOriginalColor := y.Color
	var x *Node

	if z.Left == tree.NIL {
		x = z.Right
		tree.transplant(z, z.Right)
	} else if z.Right == tree.NIL {
		x = z.Left
		tree.transplant(z, z.Left)
	} else {
		y = tree.minimum(z.Right)
		yOriginalColor = y.Color
		x = y.Right

		if y != z.Right {
			tree.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.P = y
		} else {
			x.P = y
		}

		tree.transplant(z, y)
		y.Left = z.Left
		y.Left.P = y
		y.Color = z.Color
	}

	if yOriginalColor == BLACK {
		tree.deleteFixup(x)
	}
}

func (tree *RedBlackTree) deleteFixup(x *Node) {
	for x != tree.Root && x.Color == BLACK {
		if x == x.P.Left {
			w := x.P.Right
			if w.Color == RED {
				w.Color = BLACK
				x.P.Color = RED
				tree.leftRotate(x.P)
				w = x.P.Right
			}
			if w.Left.Color == BLACK && w.Right.Color == BLACK {
				w.Color = RED
				x = x.P
			} else {
				if w.Right.Color == BLACK {
					w.Left.Color = BLACK
					w.Color = RED
					tree.rightRotate(w)
					w = x.P.Right
				}
				w.Color = x.P.Color
				x.P.Color = BLACK
				w.Right.Color = BLACK
				tree.leftRotate(x.P)
				x = tree.Root
			}
		} else {
			w := x.P.Left
			if w.Color == RED {
				w.Color = BLACK
				x.P.Color = RED
				tree.rightRotate(x.P)
				w = x.P.Left
			}
			if w.Right.Color == BLACK && w.Left.Color == BLACK {
				w.Color = RED
				x = x.P
			} else {
				if w.Left.Color == BLACK {
					w.Right.Color = BLACK
					w.Color = RED
					tree.leftRotate(w)
					w = x.P.Left
				}
				w.Color = x.P.Color
				x.P.Color = BLACK
				w.Left.Color = BLACK
				tree.rightRotate(x.P)
				x = tree.Root
			}
		}
	}

	x.Color = BLACK
}

func (tree *RedBlackTree) minimum(node *Node) *Node {
	curr := node

	for curr.Left != tree.NIL {
		curr = curr.Left
	}

	return curr
}

// replace subtree `u` with `v`
func (tree *RedBlackTree) transplant(u, v *Node) {
	if u.P == tree.NIL {
		tree.Root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}

	v.P = u.P
}
