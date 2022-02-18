package queue

import "errors"

var errOutOfRange = errors.New("trying to dequeue an empty queue")

func NewQueue() SliceQueue {
	return SliceQueue{}
}

type Queue interface {
	Enqueue(func())
	Dequeue() (func(), error)
}

type SliceQueue struct {
	length int
	queue  []func()
}

func (q *SliceQueue) Enqueue(element func()) {
	q.queue = append(q.queue, element)

	q.length++
}

func (q *SliceQueue) Dequeue() (element func(), err error) {
	if q.length < 1 {
		return nil, errOutOfRange
	}

	element = q.queue[0]

	q.queue = q.queue[1:]

	q.length--

	return element, nil
}

func (q SliceQueue) GetLength() (length int) {
	return q.length
}

func (q SliceQueue) IsEmpty() bool {
	return q.length < 1
}
