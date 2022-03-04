package queue

import "errors"

var errOutOfRange = errors.New("trying to dequeue an empty queue")

type Queue interface {
	Enqueue(interface{})
	Dequeue() (interface{}, error)
}

type SliceQueue struct {
	length int
	queue  []interface{}
}

func NewQueue() SliceQueue {
	return SliceQueue{
		length: 0,
		queue:  nil,
	}
}

func (q *SliceQueue) Enqueue(element interface{}) {
	q.queue = append(q.queue, element)

	q.length++
}

func (q *SliceQueue) Dequeue() (element interface{}, err error) {
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
