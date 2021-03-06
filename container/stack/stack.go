package stack

type Stack interface {
	Push(value interface{})
	Pop() interface{}
	Peek() interface{}
	Len() int
	IsEmpty() bool
}
