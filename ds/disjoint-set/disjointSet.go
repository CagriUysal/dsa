package disjointset

type Element[T comparable] struct {
	Value  T
	Parent *Element[T]
	Rank   int
}

type DisjointSet[T comparable] struct {
	Elements map[T]*Element[T]
}

func NewDisjointSet[T comparable]() *DisjointSet[T] {
	return &DisjointSet[T]{
		Elements: make(map[T]*Element[T]),
	}
}

func (ds *DisjointSet[T]) MakeSet(value T) {
	if _, exists := ds.Elements[value]; exists {
		return
	}

	element := &Element[T]{
		Value: value,
		Rank:  0,
	}
	element.Parent = element

	ds.Elements[value] = element
}

func (ds *DisjointSet[T]) FindSet(value T) *Element[T] {
	element, exists := ds.Elements[value]
	if !exists {
		return nil
	}

	if element != element.Parent {
		// path compression
		element.Parent = ds.FindSet(element.Parent.Value)
	}

	return element.Parent
}

func (ds *DisjointSet[T]) Union(x, y T) {
	root1 := ds.FindSet(x)
	root2 := ds.FindSet(y)

	if root1 == nil || root2 == nil {
		return
	}

	if root1 == root2 {
		return
	}

	if root1.Rank > root2.Rank {
		root2.Parent = root1
	} else if root2.Rank > root1.Rank {
		root1.Parent = root2
	} else {
		root2.Parent = root1
		root1.Rank++
	}
}
