package stack

import "errors"

var errOutOfRange = errors.New("trying to get element from an empty stack")

type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}

type SliceStack struct {
	length int
	stack  []interface{}
}

func NewStack() SliceStack {
	return SliceStack{
		length: 0,
		stack:  nil,
	}
}

func (s *SliceStack) Push(element interface{}) {
	s.stack = append(s.stack, element)

	s.length++
}

func (s *SliceStack) Pop() (element interface{}, err error) {
	if s.length < 1 {
		return nil, errOutOfRange
	}

	s.length--

	element = s.stack[s.length]

	s.stack = s.stack[:s.length]

	return element, nil
}

func (s SliceStack) Peek() (element interface{}, err error) {
	if s.length < 1 {
		return nil, errOutOfRange
	}

	element = s.stack[s.length-1]

	return element, nil
}

func (s SliceStack) GetLength() (length int) {
	return s.length
}

func (s SliceStack) IsEmpty() bool {
	return s.length < 1
}
