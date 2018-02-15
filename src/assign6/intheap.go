package main

// IntHeap is a min (max) heap of ints.
type IntHeap struct {
	hp    []int
	ismax bool
}

func (h IntHeap) Len() int      { return len(h.hp) }
func (h IntHeap) Swap(i, j int) { h.hp[i], h.hp[j] = h.hp[j], h.hp[i] }

func (h IntHeap) Less(i, j int) bool {
	if h.ismax {
		return h.hp[i] > h.hp[j]
	}
	return h.hp[i] < h.hp[j]
}

// Push push the int to the heap
func (h *IntHeap) Push(x interface{}) {
	h.hp = append(h.hp, x.(int))
}

// Pop pop the min (max) from the heap
func (h *IntHeap) Pop() interface{} {
	n := len(h.hp)
	x := h.hp[n-1]
	h.hp = h.hp[0 : n-1]
	return x
}
