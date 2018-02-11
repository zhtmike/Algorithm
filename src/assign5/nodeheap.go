package main

import "container/heap"

// Node is a data structure with id and distance
// index keeps track the position of the Node in the Heap, which is a slice of nodes.
type Node struct {
	id       int
	distance int
}

// ============================================================

// PriorityQueue implements heap.Interface hp
// The indices keeps track the position of the element in the heap
type PriorityQueue struct {
	hp      []*Node
	indices []int
}

// Len get the number of Nodes in Heap
func (h PriorityQueue) Len() int {
	return len(h.hp)
}

// Less compares the distance between two Nodes
func (h PriorityQueue) Less(i, j int) bool {
	return h.hp[i].distance < h.hp[j].distance
}

// Swap swaps between two nodes
func (h *PriorityQueue) Swap(i, j int) {
	h.hp[i], h.hp[j] = h.hp[j], h.hp[i]
	h.indices[h.hp[i].id] = i
	h.indices[h.hp[j].id] = j
}

// Push insert a node into the Heap
func (h *PriorityQueue) Push(x interface{}) {
	panic("Not implemented.")
}

// Pop pop the node with min distance
func (h *PriorityQueue) Pop() interface{} {
	n := len(h.hp)
	x := h.hp[n-1]
	h.indices[x.id] = -1
	h.hp = h.hp[:n-1]
	return x
}

// Update modify the distance and keep the structure invariant
func (h *PriorityQueue) Update(index int, distance int) {
	x := h.hp[h.indices[index]]
	x.distance = distance
	heap.Fix(h, h.indices[index])
}
