package graph

import (
	"dsa/ds/disjoint-set"
	"fmt"
	"sort"
)

// Finds minimum spanning tree in a weighted connected graph.
func MstKruskal(g WeightedGraph) []Edge {
	ds := disjointset.NewDisjointSet[string]()
	for vertex := range g.AdjList {
		ds.MakeSet(vertex)
	}

	allEdges := []Edge{}
	visitedEdges := make(map[string]struct{}) // To avoid adding duplicate edges (u,v) and (v,u)
	for _, edges := range g.AdjList {
		for _, edge := range edges {
			v1 := edge.From
			v2 := edge.To
			// This ensures ("A", "B") and ("B", "A") have the same key ("A-B").
			if v1 > v2 {
				v1, v2 = v2, v1
			}
			key := fmt.Sprintf("%s-%s", v1, v2)

			if _, visited := visitedEdges[key]; visited {
				continue
			}

			allEdges = append(allEdges, edge)
			visitedEdges[key] = struct{}{}
		}
	}

	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i].Weight < allEdges[j].Weight
	})

	mstEdges := []Edge{}
	for _, edge := range allEdges {
		if ds.FindSet(edge.From) != ds.FindSet(edge.To) {
			mstEdges = append(mstEdges, edge)
			ds.Union(edge.From, edge.To)
		}
	}

	return mstEdges
}
