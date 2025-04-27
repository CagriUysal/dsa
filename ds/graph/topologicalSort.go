package graph

import "fmt"

type DirectedGraph struct {
	AdjList map[int][]int
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{
		AdjList: make(map[int][]int),
	}
}

// AddVertex adds a vertex to the graph. If the vertex already exists,
// this is a no-op.
func (g *DirectedGraph) AddVertex(v int) {
	if _, exists := g.AdjList[v]; exists {
		return
	}

	g.AdjList[v] = []int{}
}

// AddEdge adds a directed edge from vertex u to vertex v.
// It automatically adds the vertices if they don't exist.
func (g *DirectedGraph) AddEdge(u, v int) {
	g.AddVertex(u)
	g.AddVertex(v)

	g.AdjList[u] = append(g.AdjList[u], v)
}

func (g *DirectedGraph) TopologicalSort() ([]int, error) {
	//  0 = unvisited, 1 = visiting (in current recursion stack), 2 = visited (finished)
	visited := make(map[int]int)
	var sortedOrder []int

	for vertex := range g.AdjList {
		if visited[vertex] != 0 {
			continue
		}

		if g.dfsSort(vertex, visited, &sortedOrder) {
			return nil, fmt.Errorf("graph contains a cycle, topological sort not possible")
		}
	}

	// The sortedOrder is built in reverse order in dfsSort, so reverse it
	// to get the correct topological order.
	for i, j := 0, len(sortedOrder)-1; i < j; i, j = i+1, j-1 {
		sortedOrder[i], sortedOrder[j] = sortedOrder[j], sortedOrder[i]
	}

	return sortedOrder, nil
}

// dfsSort is a helper function for the recursive DFS-based topological sort.
// It returns true if a cycle is detected during traversal, false otherwise.
func (g *DirectedGraph) dfsSort(v int, visited map[int]int, sortedOrder *[]int) bool {
	visited[v] = 1

	for _, neighbor := range g.AdjList[v] {
		// cycle detected
		if visited[neighbor] == 1 {
			return true
		}

		if visited[neighbor] == 2 {
			continue
		}

		if g.dfsSort(neighbor, visited, sortedOrder) {
			return true
		}
	}

	visited[v] = 2
	*sortedOrder = append(*sortedOrder, v)
	return false
}
