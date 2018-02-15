package main

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	var inthp IntHeap
	items := []int{3, 2, 1, 4, 5}
	inthp.hp = items
	heap.Init(&inthp)
	x := heap.Pop(&inthp).(int)
	if x != 1 {
		t.Error("Expected 1 got ", x)
	}
	heap.Push(&inthp, -1)
	x = heap.Pop(&inthp).(int)
	if x != -1 {
		t.Error("Expected -1 got ", x)
	}
}

func TestMaxHeap(t *testing.T) {
	var inthp IntHeap
	items := []int{3, 2, 1, 4, 5}
	inthp.hp = items
	inthp.ismax = true
	heap.Init(&inthp)
	x := heap.Pop(&inthp).(int)
	if x != 5 {
		t.Error("Expected 5 got ", x)
	}
	heap.Push(&inthp, -1)
	x = heap.Pop(&inthp).(int)
	if x != 4 {
		t.Error("Expected 4 got ", x)
	}
}
