package algorithms

import "../structures"

// Stack implements a stack by structures
type Stack struct {
	*structures.LinkList
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{structures.NewLinkList()}
}

// Push ...
func (s *Stack) Push(item interface{}) {
	s.Add(item)
}

// Pop ...
func (s *Stack) Pop() (item interface{}) {

	if s.IsEmpty() {
		panic("Stack Empty")
	}

	item = s.Del()
	return
}
