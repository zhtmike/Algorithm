package main

import (
	"container/heap"
	"utils/graph"
)

// DijkstraAlgo returns the sortest path in a weigted adjacent list
func DijkstraAlgo(g graph.WAdjlist) []int {
	// Initialize the heap (PriorityQueue), the order is not preserved
	var pq PriorityQueue
	pq.hp = make([]*Node, len(g))
	pq.indices = make([]int, len(g))
	for i := 0; i < len(g); i++ {
		pq.indices[i] = i
		pq.hp[i] = &Node{
			id:       i,
			distance: int(10E9),
		}
	}
	pq.hp[0].distance = 0
	heap.Init(&pq)

	// Initialize the shortest distance
	shortD := make([]int, len(g))

	// Start the search with index 0
	nodeI := pq.hp[pq.indices[0]]
	shortD[nodeI.id] = nodeI.distance
	for pq.Len() > 0 {
		for i := 0; i < len(g[nodeI.id]); i++ {
			// compute the shortest distance between node i and its neighborhood node j
			nodeJind := pq.indices[g[nodeI.id][i].Vertex]
			// If nodeJ's index is -1, then it means that it is already explored
			if nodeJind >= 0 {
				nodeJ := pq.hp[nodeJind]
				distance := min(nodeJ.distance, nodeI.distance+g[nodeI.id][i].Weight)
				pq.Update(g[nodeI.id][i].Vertex, distance)
			}
		}
		nodeI = heap.Pop(&pq).(*Node)
		shortD[nodeI.id] = nodeI.distance
	}

	return shortD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
