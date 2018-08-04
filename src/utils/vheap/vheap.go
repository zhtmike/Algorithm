package vheap

import "container/heap"

// Node is a data structure with id and distance
// index keeps track the position of the Node in the Heap, which is a slice of nodes.
type Node struct {
	ID       int
	Distance int
}

// ============================================================

// PriorityQueue implements heap.Interface hp
// The indices keeps track the position of the element in the heap
type PriorityQueue struct {
	Hp      []*Node
	Indices []int
}

// Len get the number of Nodes in Heap
func (h PriorityQueue) Len() int {
	return len(h.Hp)
}

// Less compares the distance between two Nodes
func (h PriorityQueue) Less(i, j int) bool {
	return h.Hp[i].Distance < h.Hp[j].Distance
}

// Swap swaps between two nodes
func (h *PriorityQueue) Swap(i, j int) {
	h.Hp[i], h.Hp[j] = h.Hp[j], h.Hp[i]
	h.Indices[h.Hp[i].ID] = i
	h.Indices[h.Hp[j].ID] = j
}

// Push insert a node into the Heap
func (h *PriorityQueue) Push(x interface{}) {
	panic("Not implemented.")
}

// Pop pop the node with min distance
func (h *PriorityQueue) Pop() interface{} {
	n := len(h.Hp)
	x := h.Hp[n-1]
	h.Indices[x.ID] = -1
	h.Hp = h.Hp[:n-1]
	return x
}

// Update modify the distance and keep the structure invariant
func (h *PriorityQueue) Update(index int, distance int) {
	x := h.Hp[h.Indices[index]]
	x.Distance = distance
	heap.Fix(h, h.Indices[index])
}
