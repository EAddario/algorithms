package structures

// Stack implements a stack by structures
type Stack struct {
	*LinkList
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{NewLinkList()}
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
