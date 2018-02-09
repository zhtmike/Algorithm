package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	var stack Stack
	stack.Push(0)
	if stack.top.value != 0 {
		t.Error("Expected 0 got ", stack.top.value)
	}
	stack.Push(1)
	if stack.top.value != 1 {
		t.Error("Expected 1 got ", stack.top.value)
	}
}

func TestPop(t *testing.T) {
	var stack Stack
	stack.Push(0)
	stack.Push(1)
	v := stack.Pop()
	if v != 1 {
		t.Error("Expected 1 got ", v)
	}
	v = stack.Pop()
	if v != 0 {
		t.Error("Expected 0 got ", v)
	}
	v = stack.Pop()
	if v != nil {
		t.Error("Expected nil got ", v)
	}
}

func TestLen(t *testing.T) {
	var stack Stack
	stack.Push(0)
	stack.Push(1)
	if stack.Len() != 2 {
		t.Error("Expected 2 got ", stack.Len())
	}
	stack.Pop()
	if stack.Len() != 1 {
		t.Error("Expected 1 got ", stack.Len())
	}
	stack.Pop()
	if stack.Len() != 0 {
		t.Error("Expected 0 got ", stack.Len())
	}
}
