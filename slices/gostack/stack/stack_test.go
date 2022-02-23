package stack_test

import (
	"testing"

	"github.com/nndergunov/basicGo/slices/gostack/stack"
)

func TestPushPop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
	}{
		{"bool push/pop", []interface{}{true, true, false, true, false}},
		{"string push/pop", []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"}},
		{"int push/pop", []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7}},
		{"rune push/pop", []interface{}{'i', 'a', 'i', '4', 'p', 'q'}},
		{"float64 push/pop", []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			s := stack.NewStack()

			for i := 0; i < len(test.elements); i++ {
				s.Push(test.elements[i])
			}

			for i := 0; i < len(test.elements); i++ {
				el, err := s.Pop()
				if err != nil {
					t.Errorf("pop error: %v", err)
				}

				currentElement := test.elements[len(test.elements)-i-1]

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
					t.Errorf("got unexpected type")
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

			s := stack.NewStack()

			if !s.IsEmpty() {
				t.Error("stack should be empty, but is not")
			}

			for i := 0; i < len(test.elements); i++ {
				expStackLength := i + 1

				s.Push(test.elements[i])

				if s.GetLength() != expStackLength {
					t.Errorf("expected %d values in stack, got %d", expStackLength, s.GetLength())
				}
			}

			for i := 0; i < len(test.elements); i++ {
				expStackLength := len(test.elements) - i - 1

				_, err := s.Pop()
				if err != nil {
					t.Errorf("pop error: %v", err)
				}

				if s.GetLength() != expStackLength {
					t.Errorf("expected %d values in stack, got %d", expStackLength, s.GetLength())
				}
			}
		})
	}
}
