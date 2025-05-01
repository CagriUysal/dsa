package graph_test

import (
	"dsa/ds/graph"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAdjListGraph_BFS(t *testing.T) {
	tests := []struct {
		name          string
		graph         *graph.UnweightedGraph
		startNode     int
		wantDistances map[int]int
	}{
		{
			name: "Simple connected graph",
			graph: func() *graph.UnweightedGraph {
				g := graph.NewGraph()
				g.AddEdge(0, 1)
				g.AddEdge(0, 2)
				g.AddEdge(1, 3)
				g.AddEdge(2, 4)
				g.AddEdge(3, 4)
				return g
			}(),
			startNode: 0,
			wantDistances: map[int]int{
				0: 0,
				1: 1,
				2: 1,
				3: 2,
				4: 2,
			},
		},
		{
			name: "Graph with disconnected component",
			graph: func() *graph.UnweightedGraph {
				g := graph.NewGraph()
				g.AddEdge(0, 1)
				g.AddEdge(2, 3)
				return g
			}(),
			startNode: 0,
			wantDistances: map[int]int{
				0: 0,
				1: 1,
			},
		},
		{
			name: "Graph with cycle",
			graph: func() *graph.UnweightedGraph {
				g := graph.NewGraph()
				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 0) // Cycle 0-1-2
				g.AddEdge(2, 3)
				return g
			}(),
			startNode: 0,
			wantDistances: map[int]int{
				0: 0,
				1: 1,
				2: 1,
				3: 2,
			},
		},
		{
			name: "Single node graph",
			graph: func() *graph.UnweightedGraph {
				g := graph.NewGraph()
				g.AddVertex(0)
				return g
			}(),
			startNode:     0,
			wantDistances: map[int]int{0: 0},
		},
		{
			name:          "Empty graph",
			graph:         graph.NewGraph(),
			startNode:     0,
			wantDistances: map[int]int{0: 0},
		},
		{
			name: "Linear graph",
			graph: func() *graph.UnweightedGraph {
				g := graph.NewGraph()
				g.AddEdge(0, 1)
				g.AddEdge(1, 2)
				g.AddEdge(2, 3)
				g.AddEdge(3, 4)
				return g
			}(),
			startNode: 0,
			wantDistances: map[int]int{
				0: 0,
				1: 1,
				2: 2,
				3: 3,
				4: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.graph
			gotDistances := graph.BFS(g, tt.startNode)
			// Use cmpopts.EquateEmpty() because an empty map and a nil map should be considered equal in tests
			// However, BFS initializes the map, so it will be empty, not nil.
			// Let's compare maps directly. If wantDistances is nil (e.g., empty graph test), create an empty map.
			expected := tt.wantDistances
			if expected == nil {
				expected = make(map[int]int)
			}

			if diff := cmp.Diff(expected, gotDistances); diff != "" {
				t.Errorf("BFS() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
