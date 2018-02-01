/*
Package stack implement stack by linked list
Ref: https://gist.github.com/bemasher/1777766 */
package stack

// Stack is a data structure with FILO mechanism
type Stack struct {
	top  *Element
	size int
}

// Element can be any types satisfy the empty interface
type Element struct {
	value interface{}
	next  *Element
}

// Len return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Pop remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}
