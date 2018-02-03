package assign5

import (
	"testing"
)

func TestHeapInit(t *testing.T) {
	var hp Heap
	var nodes []*Node
	hp.Init(nodes)
	if len(*hp.pq) != 0 {
		t.Error("Expected 0 got ", len(*hp.pq))
	}
	nodes = make([]*Node, 1)
	nodes[0] = &Node{
		id:       1,
		distance: 1,
	}
	hp.Init(nodes)
	if len(*hp.pq) != 1 {
		t.Error("Expected 1 got ", len(*hp.pq))
	}
}

func TestHeapPushPop(t *testing.T) {
	var hp Heap
	items := map[int]int{
		1: 999, 2: 200, 3: 400,
	}
	nodes := make([]*Node, len(items))
	i := 0
	for value, priority := range items {
		nodes[i] = &Node{
			id:       value,
			distance: priority,
		}
		i++
	}
	hp.Push(nodes[0])
	hp.Push(nodes[1])
	hp.Push(nodes[2])
	node := hp.Pop()
	if node.id != 2 {
		t.Error("Expected 2 got ", node.id)
	}
	node = hp.Pop()
	if node.id != 3 {
		t.Error("Expected 3 got ", node.id)
	}
	node = hp.Pop()
	if node.id != 1 {
		t.Error("Expected 1 got ", node.id)
	}
	node = hp.Pop()
	if node != nil {
		t.Error("Expected nil got ", node)
	}
}

func TestHeapPushPopWithUpdate(t *testing.T) {
	var hp Heap
	items := map[int]int{
		1: 999, 2: 200, 3: 400,
	}
	nodes := make([]*Node, len(items))
	i := 0
	for value, priority := range items {
		nodes[i] = &Node{
			id:       value,
			distance: priority,
		}
		i++
	}
	hp.Push(nodes[0])
	hp.Push(nodes[1])
	hp.Push(nodes[2])
	hp.Update(nodes[0], 100)
	node := hp.Pop()
	if node.id != 1 {
		t.Error("Expected 1 got ", node.id)
	}
	node = hp.Pop()
	if node.id != 2 {
		t.Error("Expected 2 got ", node.id)
	}
	node = hp.Pop()
	if node.id != 3 {
		t.Error("Expected 3 got ", node.id)
	}
	node = hp.Pop()
	if node != nil {
		t.Error("Expected nil got ", node)
	}
}
