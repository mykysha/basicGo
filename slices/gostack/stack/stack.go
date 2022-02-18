package stack

import "errors"

var errOutOfRange = errors.New("trying to get element from an empty stack")

func NewStack() SliceStack {
	return SliceStack{}
}

type Stack interface {
	Push(func())
	Pop() (func(), error)
	Peek() (func(), error)
}

type SliceStack struct {
	length int
	stack  []func()
}

func (s *SliceStack) Push(element func()) {
	s.stack = append(s.stack, element)

	s.length++
}

func (s *SliceStack) Pop() (element func(), err error) {
	if s.length < 1 {
		return nil, errOutOfRange
	}

	s.length--

	element = s.stack[s.length]

	s.stack = s.stack[:s.length]

	return element, nil
}

func (s SliceStack) Peek() (element func(), err error) {
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
