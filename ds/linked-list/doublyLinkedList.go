package linkedlist

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	key  int
	next *Node
	prev *Node
}

func (l *DoublyLinkedList) Search(k int) *Node {
	current := l.head
	for current != nil && current.key != k {
		current = current.next
	}

	return current
}

func (l *DoublyLinkedList) Prepend(value int) {
	newNode := &Node{
		key: value,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

// inserts x after the y
func (l *DoublyLinkedList) Insert(x, y *Node) {
	if y == nil || x == nil {
		return
	}

	x.prev = y
	x.next = y.next

	if y.next != nil {
		y.next.prev = x
	}
	y.next = x

	if x.next == nil {
		l.tail = x
	}
}

func (l *DoublyLinkedList) Delete(x *Node) {
	if x == nil {
		return
	}

	if x.prev != nil {
		x.prev.next = x.next
	} else {
		l.head = x.next
	}

	if x.next != nil {
		x.next.prev = x.prev
	} else {
		l.tail = x.prev
	}
}
