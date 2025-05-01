package graph

import "math"

// DagShortestPaths calculates the shortest paths from a single source vertex
// in a directed acyclic weighted graph(DAG). It can handle negative weights.
//
// Args:
//
//	g: The weighted graph.
//	startVertex: The starting vertex for calculating shortest paths.
//
// Returns:
//
//	*ShortestPath: A pointer to a ShortestPath struct containing distances and parents,
//	or nil if a negative cycle is detected.
func DagShortestPaths(g WeightedGraph, startVertex string) *ShortestPath {
	topologicalOrder, err := TopologicalSort(g)
	if err != nil {
		panic(err)
	}

	shortestPath := &ShortestPath{
		distances: make(map[string]int),
		parents:   make(map[string]*string),
	}
	
	for  vertex := range 	g.AdjList {
			shortestPath.distances[vertex] = int(math.Inf(1))
			shortestPath.parents[vertex] = nil
	}		
	shortestPath.distances[startVertex] = 0

	for _, vertex := range topologicalOrder {
			if shortestPath.distances[vertex] == int(math.Inf(1)){
					continue
			}

			for _, edge := range g.AdjList[vertex] {
				newDist := shortestPath.distances[vertex] + edge.Weight
				if newDist < shortestPath.distances[edge.To] {
					shortestPath.distances[edge.To] = newDist
					currentParent := vertex
					shortestPath.parents[edge.To] = &currentParent
				}
			}
	}

	return shortestPath
}
