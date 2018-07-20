package stack

import . "github.com/Marsyyh/go-datastructure/container/linkedlist"

type LinkedListStack struct {
	top    *Node
	length int
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil, 0}
}

func (s *LinkedListStack) Push(value interface{}) {
	n := &Node{value, s.top}
	s.top = n
	s.length++
}

func (s *LinkedListStack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	value := s.top.Value
	s.top = s.top.Next
	s.length--
	return value
}

func (s *LinkedListStack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.Value
}

func (s *LinkedListStack) Len() int {
	return s.length
}

func (s *LinkedListStack) IsEmpty() bool {
	return s.top == nil
}
