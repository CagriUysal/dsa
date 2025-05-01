package graph

// UnweightedGraph represents an unweighted, undirected graph using an adjacency list.
type UnweightedGraph struct {
	AdjList map[int][]int
}

func NewGraph() *UnweightedGraph {
	return &UnweightedGraph{
		AdjList: make(map[int][]int),
	}
}

func (g *UnweightedGraph) AddVertex(v int) {
	if _, exists := g.AdjList[v]; !exists {
		g.AdjList[v] = []int{}
	}
}

func (g *UnweightedGraph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)

	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}

type Edge struct {
	From   string
	To     string
	Weight int
}

type WeightedGraph struct {
	AdjList map[string][]Edge
}

func (g *WeightedGraph) AddVertex(v string) {
	if g.AdjList == nil {
		g.AdjList = make(map[string][]Edge)
	}

	if _, exists := g.AdjList[v]; !exists {
		g.AdjList[v] = []Edge{}
	}
}

func (g *WeightedGraph) AddEdge(from string, to string, weight int) {
	if g.AdjList == nil {
		g.AdjList = make(map[string][]Edge)
	}

	g.AddVertex(from)
	g.AddVertex(to)

	g.AdjList[from] = append(g.AdjList[from], Edge{From: from, To: to, Weight: weight})
	g.AdjList[to] = append(g.AdjList[to], Edge{From: to, To: from, Weight: weight})
}
