package nodeheap

import (
	"container/heap"
)

// Node is a data structure with id and distance
// index keeps track the position of the Node in the Heap, which is a slice of nodes.
type Node struct {
	id       int
	distance int
	index    int
}

// ============================================================

// PriorityQueue implements heap.Interface and holds Items.
type priorityQueue []*Node

// Len get the number of Nodes in Heap
func (h priorityQueue) Len() int {
	return len(h)
}

// Less compares the distance between two Nodes
func (h priorityQueue) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

// Swap swaps between two nodes
func (h priorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

// Push insert a node into the Heap
func (h *priorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.index = len(*h)
	*h = append(*h, node)
}

// Pop pop the node with min distance
func (h *priorityQueue) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	x.index = -1
	*h = (*h)[:n-1]
	return x
}

// ============================================================

// Heap is a data structure supports O(1) pop min
// It is a wrapper of priorityQueue for easy accessing
type Heap struct {
	pq *priorityQueue
}

// Init inits the heap with a slice of node objects
func (h *Heap) Init(nodes []*Node) {
	items := priorityQueue(nodes)
	heap.Init(&items)
	h.pq = &items
}

// Push push a node into Heap
func (h *Heap) Push(node *Node) {
	if h.pq != nil {
		heap.Push(h.pq, node)
	} else {
		tmp := make(priorityQueue, 1)
		tmp[0] = node
		heap.Init(&tmp)
		h.pq = &tmp
	}
}

// Pop pop a node from Heap
func (h *Heap) Pop() (node *Node) {
	if len(*h.pq) > 0 {
		node = heap.Pop(h.pq).(*Node)
		return
	}
	return nil
}

// Update modify the distance and keep the structure invariant
func (h *Heap) Update(x *Node, distance int) {
	x.distance = distance
	heap.Fix(h.pq, x.index)
}
