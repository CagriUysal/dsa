package graph

// runs BFS on the source `s` vertex and returns the distances for each discovered vertex
func BFS(g *UnweightedGraph, s int) map[int]int {
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
