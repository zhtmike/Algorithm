package main

import (
	"container/heap"
	"utils/graph"
	"utils/vheap"
)

// DijkstraAlgo returns the sortest path in a weigted adjacent list
func DijkstraAlgo(g graph.WAdjlist) []int {
	// Initialize the heap (PriorityQueue), the order is not preserved
	var pq vheap.PriorityQueue
	pq.Hp = make([]*vheap.Node, len(g))
	pq.Indices = make([]int, len(g))
	for i := 0; i < len(g); i++ {
		pq.Indices[i] = i
		pq.Hp[i] = &vheap.Node{
			ID:       i,
			Distance: int(10E9),
		}
	}
	pq.Hp[0].Distance = 0
	heap.Init(&pq)

	// Initialize the shortest distance
	shortD := make([]int, len(g))

	// Start the search with index 0
	nodeI := pq.Hp[pq.Indices[0]]
	shortD[nodeI.ID] = nodeI.Distance
	for pq.Len() > 0 {
		for i := 0; i < len(g[nodeI.ID]); i++ {
			// compute the shortest distance between node i and its neighborhood node j
			nodeJind := pq.Indices[g[nodeI.ID][i].Vertex]
			// If nodeJ's index is -1, then it means that it is already explored
			if nodeJind >= 0 {
				nodeJ := pq.Hp[nodeJind]
				distance := min(nodeJ.Distance, nodeI.Distance+g[nodeI.ID][i].Weight)
				pq.Update(g[nodeI.ID][i].Vertex, distance)
			}
		}
		nodeI = heap.Pop(&pq).(*vheap.Node)
		shortD[nodeI.ID] = nodeI.Distance
	}

	return shortD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
