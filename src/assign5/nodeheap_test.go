package main

import (
	"container/heap"
	"testing"
)

func TestHeap(t *testing.T) {
	items := map[int]int{
		0: 999, 1: 200, 2: 400,
	}

	var pq PriorityQueue
	pq.hp = make([]*Node, len(items))
	pq.indices = make([]int, len(items))
	for i := 0; i < len(items); i++ {
		pq.indices[i] = i
		pq.hp[i] = &Node{
			id:       i,
			distance: items[i],
		}
	}
	heap.Init(&pq)
	pq.Update(2, 100)
	x := heap.Pop(&pq).(*Node)
	if x.id != 2 {
		t.Error("Expected 2 got ", x.id)
	}
	pq.Update(0, 100)
	x = heap.Pop(&pq).(*Node)
	if x.id != 0 {
		t.Error("Expected 0 got ", x.id)
	}
	x = heap.Pop(&pq).(*Node)
	if x.id != 1 {
		t.Error("Expected 1 got ", x.id)
	}
}
