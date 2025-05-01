package graph

import (
	"container/heap"
)

// PriorityQueue implements the heap.Interface for a min-heap of Edges,
// ordered by edge weight.
type priorityQue []*Edge

func (pq priorityQue) Len() int {
	return len(pq)
}

func (pq priorityQue) Less(i, j int) bool {
	return pq[i].Weight < pq[j].Weight
}

func (pq priorityQue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQue) Push(x any) {
	item := x.(*Edge)
	*pq = append(*pq, item)
}

func (pq *priorityQue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func MstPrims(g WeightedGraph, startVertex string) []Edge {
	mstEdges := []Edge{}
	visited := make(map[string]struct{})

	pq := make(priorityQue, 0)
	heap.Init(&pq)

	visited[startVertex] = struct{}{}

	for _, edge := range g.AdjList[startVertex] {
		heap.Push(&pq, &edge)
	}

	for len(pq) > 0 {
		minEdge := heap.Pop(&pq).(*Edge)

		if _, exists := visited[minEdge.To]; exists {
			continue
		}

		visited[minEdge.To] = struct{}{}
		mstEdges = append(mstEdges, *minEdge)

		for _, edge := range g.AdjList[minEdge.To] {
			if _, exists := visited[edge.To]; exists {
				continue
			}

			heap.Push(&pq, &edge)
		}
	}

	return mstEdges
}
