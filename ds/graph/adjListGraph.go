package graph

// AdjListGraph represents an unweighted, undirected graph using an adjacency list.
type AdjListGraph struct {
	AdjList map[int][]int
}

func NewGraph() *AdjListGraph {
	return &AdjListGraph{
		AdjList: make(map[int][]int),
	}
}

func (g *AdjListGraph) AddVertex(v int) {
	if _, exists := g.AdjList[v]; !exists {
		g.AdjList[v] = []int{}
	}
}

func (g *AdjListGraph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)

	g.AdjList[u] = append(g.AdjList[u], v)
	g.AdjList[v] = append(g.AdjList[v], u)
}
