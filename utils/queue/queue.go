/*
Package queue implement queue by linked list.
*/
package queue

// Queue is a data structure with FIFO mechanism
type Queue struct {
	first *Element
	last  *Element
	size  int
}

// Element can be any types satisfy the empty interface
type Element struct {
	value interface{}
	next  *Element
}

// Len return the queue's length
func (s *Queue) Len() int {
	return s.size
}

// Push a new element into the queue
func (s *Queue) Push(value interface{}) {
	if s.size > 0 {
		s.last.next = &Element{value, nil}
		s.last = s.last.next
	} else {
		s.first = &Element{value, nil}
		s.last = s.first
	}
	s.size++
}

// Pop remove the first element from the queue and return it's value
// If the queue is empty, return nil
func (s *Queue) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.first = s.first.value, s.first.next
		s.size--
		return
	}
	return nil
}
