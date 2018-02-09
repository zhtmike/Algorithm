package queue

import (
	"testing"
)

func TestPush(t *testing.T) {
	var queue Queue
	queue.Push(0)
	if queue.first.value != 0 {
		t.Error("Expected 0 got ", queue.first.value)
	}
	if queue.last.value != 0 {
		t.Error("Expected 0 got ", queue.last.value)
	}
	queue.Push(1)
	if queue.first.value != 0 {
		t.Error("Expected 0 got ", queue.first.value)
	}
	if queue.last.value != 1 {
		t.Error("Expected 1 got ", queue.last.value)
	}
}

func TestPop(t *testing.T) {
	var queue Queue
	queue.Push(0)
	queue.Push(1)
	v := queue.Pop()
	if v != 0 {
		t.Error("Expected 0 got ", v)
	}
	v = queue.Pop()
	if v != 1 {
		t.Error("Expected 1 got ", v)
	}
	v = queue.Pop()
	if v != nil {
		t.Error("Expected nil got ", v)
	}
}

func TestLen(t *testing.T) {
	var queue Queue
	queue.Push(0)
	queue.Push(1)
	if queue.Len() != 2 {
		t.Error("Expected 2 got ", queue.Len())
	}
	queue.Pop()
	if queue.Len() != 1 {
		t.Error("Expected 1 got ", queue.Len())
	}
	queue.Pop()
	if queue.Len() != 0 {
		t.Error("Expected 0 got ", queue.Len())
	}
}
