package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		next  *node
	}
)

func New() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	value := s.top.value
	s.top = s.top.next
	s.length--
	return value
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
