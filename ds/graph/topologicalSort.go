package graph

import "fmt"

func TopologicalSort(g WeightedGraph) ([]string, error) {
	//  0 = unvisited, 1 = visiting (in current recursion stack), 2 = visited (finished)
	visited := make(map[string]int)
	var sortedOrder []string

	for vertex := range g.AdjList {
		if visited[vertex] != 0 {
			continue
		}

		if dfsSort(g, vertex, visited, &sortedOrder) {
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
func dfsSort(g WeightedGraph, v string, visited map[string]int, sortedOrder *[]string) bool {
	visited[v] = 1

	for _, edge := range g.AdjList[v] {
		// cycle detected
		neighbor := edge.To
		if visited[neighbor] == 1 {
			return true
		}

		if visited[neighbor] == 2 {
			continue
		}

		if dfsSort(g, neighbor, visited, sortedOrder) {
			return true
		}
	}

	visited[v] = 2
	*sortedOrder = append(*sortedOrder, v)
	return false
}
