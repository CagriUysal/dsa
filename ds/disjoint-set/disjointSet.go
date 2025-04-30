package disjointset

type Element struct {
	Value  int
	Parent *Element
	Rank   int
}

type DisjointSet struct {
	Elements map[int]*Element
}

func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		Elements: make(map[int]*Element),
	}
}

func (ds *DisjointSet) MakeSet(value int) {
	if _, exists := ds.Elements[value]; exists {
		return
	}

	element := &Element{
		Value: value,
		Rank:  0,
	}
	element.Parent = element

	ds.Elements[value] = element
}

func (ds *DisjointSet) FindSet(value int) *Element {
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

func (ds *DisjointSet) Union(x, y int) {
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
