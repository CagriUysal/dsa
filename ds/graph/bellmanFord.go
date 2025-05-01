package graph

import "math"

// BellmanFord calculates the shortest paths from a single source vertex
// in a weighted graph. It can handle negative edge weights and detects
// negative weight cycles reachable from the source.
//
// Args:
//
//	g: The weighted graph.
//	sourceVertex: The starting vertex for calculating shortest paths.
//
// Returns:
//
//	*Path: A pointer to a Path struct containing distances and parents,
//	       or nil if a negative cycle is detected.
//	bool: True if shortest paths are found without negative cycles,
//	      False if a negative cycle reachable from the source is detectd.
func BellmanFord(g WeightedGraph, sourceVertex string) (*ShortestPath, bool) {
	path := &ShortestPath{
		distances: make(map[string]int),
		parents:   make(map[string]*string),
	}

	for vertex := range g.AdjList {
		path.distances[vertex] = int(math.Inf(1))
		path.parents[vertex] = nil
	}
	path.distances[sourceVertex] = 0

	allEdges := []Edge{}
	for _, edges := range g.AdjList {
		for _, edge := range edges {
			allEdges = append(allEdges, edge)
		}
	}

	numVertices := len(g.AdjList)
	for range numVertices - 1 {
		for _, edge := range allEdges {
			if path.distances[edge.From] != int(math.Inf(1)) &&
				path.distances[edge.From]+edge.Weight < path.distances[edge.To] {
				path.parents[edge.To] = &edge.From
				path.distances[edge.To] = path.distances[edge.From] + edge.Weight
			}
		}
	}

	for _, edge := range allEdges {
		if path.distances[edge.To] > path.distances[edge.From]+edge.Weight {
			return nil, false
		}
	}

	return path, true
}
