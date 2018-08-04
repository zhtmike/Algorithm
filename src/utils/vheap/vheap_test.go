package vheap

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	items := map[int]int{
		0: 999, 1: 200, 2: 400,
	}

	var pq PriorityQueue
	pq.Hp = make([]*Node, len(items))
	pq.Indices = make([]int, len(items))
	for i := 0; i < len(items); i++ {
		pq.Indices[i] = i
		pq.Hp[i] = &Node{
			ID:       i,
			Distance: items[i],
		}
	}
	heap.Init(&pq)
	pq.Update(2, 100)
	x := heap.Pop(&pq).(*Node)
	if x.ID != 2 {
		t.Error("Expected 2 got ", x.ID)
	}
	pq.Update(0, 100)
	x = heap.Pop(&pq).(*Node)
	if x.ID != 0 {
		t.Error("Expected 0 got ", x.ID)
	}
	x = heap.Pop(&pq).(*Node)
	if x.ID != 1 {
		t.Error("Expected 1 got ", x.ID)
	}
}
