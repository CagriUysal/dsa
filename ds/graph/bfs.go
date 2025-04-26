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

// runs BFS on the source `s` vertex and returns the distances for each discovered vertex
func (g *AdjListGraph) BFS(s int) map[int]int {
	visited := make(map[int]bool)
	distances := make(map[int]int)

	que := []int{}

	que = append(que, s)
	visited[s] = true
	distances[s] = 0

	for len(que) > 0 {
		curr := que[0]
		que = que[1:]

		for _, neighbor := range g.AdjList[curr] {
			if visited[neighbor] {
				continue
			}

			que = append(que, neighbor)
			visited[neighbor] = true
			distances[neighbor] = distances[curr] + 1
		}
	}

	return distances
}
