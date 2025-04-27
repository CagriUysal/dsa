package graph

func (g *AdjListGraph) DFS(startVertex int) []int {
	visited := make(map[int]bool)
	visitOrder := []int{}

	g.recurseDFS(startVertex, visited, &visitOrder)

	for v := range g.AdjList {
		if visited[v] {
			continue
		}

		g.recurseDFS(v, visited, &visitOrder)
	}

	return visitOrder
}

func (g *AdjListGraph) recurseDFS(v int, visited map[int]bool, visitOrder *[]int) {
	visited[v] = true
	*visitOrder = append(*visitOrder, v)

	for _, neighboor := range g.AdjList[v] {
		if visited[neighboor] {
			continue
		}
		g.recurseDFS(neighboor, visited, visitOrder)
	}
}
