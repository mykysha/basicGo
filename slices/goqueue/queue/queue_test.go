package queue_test

import (
	"testing"

	"github.com/nndergunov/basicGo/slices/goqueue/queue"
)

func TestQueueDequeue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
	}{
		{"bool queue/dequeue", []interface{}{true, true, false, true, false}},
		{"string queue/dequeue", []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"}},
		{"int queue/dequeue", []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7}},
		{"rune queue/dequeue", []interface{}{'i', 'a', 'i', '4', 'p', 'q'}},
		{"float64 queue/dequeue", []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			q := queue.NewQueue()

			for i := 0; i < len(test.elements); i++ {
				q.Enqueue(test.elements[i])
			}

			for i := 0; i < len(test.elements); i++ {
				el, err := q.Dequeue()
				if err != nil {
					t.Errorf("dequeue error: %v", err)
				}

				currentElement := test.elements[i]

				switch exp := currentElement.(type) {
				case bool:
					got, ok := el.(bool)
					if !ok {
						t.Error("expected bool type, did not get it")
					}

					if got != exp {
						t.Errorf("expected %v, got %v", exp, got)
					}
				case string:
					got, ok := el.(string)
					if !ok {
						t.Error("expected string type, did not get it")
					}

					if got != exp {
						t.Errorf("expected %v, got %v", exp, got)
					}
				case int:
					got, ok := el.(int)
					if !ok {
						t.Error("expected int type, did not get it")
					}

					if got != exp {
						t.Errorf("expected %v, got %v", exp, got)
					}
				case rune:
					got, ok := el.(rune)
					if !ok {
						t.Error("expected rune type, did not get it")
					}

					if got != exp {
						t.Errorf("expected %v, got %v", exp, got)
					}
				case float64:
					got, ok := el.(float64)
					if !ok {
						t.Error("expected float64 type, did not get it")
					}

					if got != exp {
						t.Errorf("expected %v, got %v", exp, got)
					}
				default:
					t.Error("got unexpected type")
				}
			}
		})
	}
}

func TestElementNumber(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
	}{
		{"bool element number", []interface{}{true, true, false, true, false}},
		{"string element number", []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"}},
		{"int element number", []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7}},
		{"rune element number", []interface{}{'i', 'a', 'i', '4', 'p', 'q'}},
		{"float64 element number", []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			q := queue.NewQueue()

			if !q.IsEmpty() {
				t.Error("queue should be empty, but is not")
			}

			for i := 0; i < len(test.elements); i++ {
				expQueueLength := i + 1

				q.Enqueue(test.elements[i])

				if q.GetLength() != expQueueLength {
					t.Errorf("expected %d values in queue, got %d", expQueueLength, q.GetLength())
				}
			}

			for i := 0; i < len(test.elements); i++ {
				expQueueLength := len(test.elements) - i - 1

				_, err := q.Dequeue()
				if err != nil {
					t.Errorf("dequeue error: %v", err)
				}

				if q.GetLength() != expQueueLength {
					t.Errorf("expected %d values in queue, got %d", expQueueLength, q.GetLength())
				}
			}
		})
	}
}
