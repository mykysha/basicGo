package binarytree_test

import (
	"errors"
	"testing"

	"github.com/nndergunov/basicGo/slices/gotree/binarytree"
)

func TestBinaryTreeOrder(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
		sorted   []interface{}
	}{
		{
			testName: "bool order test",
			elements: []interface{}{true, false},
			sorted:   []interface{}{false, true},
		},
		{
			testName: "string order test",
			elements: []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"},
			sorted:   []interface{}{"apple", "dino", "phone", "pi", "tea", "uno"},
		},
		{
			testName: "int order test",
			elements: []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7},
			sorted:   []interface{}{-987, -54, 1, 7, 13, 23, 55, 65, 87, 458, 990},
		},
		{
			testName: "rune order test",
			elements: []interface{}{'i', 'a', 'm', '4', 'p', 'q'},
			sorted:   []interface{}{'4', 'a', 'i', 'm', 'p', 'q'},
		},
		{
			testName: "float64 order test",
			elements: []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7},
			sorted:   []interface{}{0.5, 2.1, 4.56789, 5.5, 6.66, 987.7},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			bt := binarytree.NewBinaryTree()

			for i := 0; i < len(test.elements); i++ {
				err := bt.Insert(test.elements[i])
				if err != nil {
					t.Errorf("did not expect error %v", err)
				}
			}

			order := bt.GetInOrder()

			for key, val := range test.sorted {
				el := order[key]

				switch exp := val.(type) {
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

func TestGetMinValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
		min      interface{}
	}{
		{"bool min value", []interface{}{true, false}, false},
		{"string min value", []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"}, "apple"},
		{"int min value", []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7}, -987},
		{"rune min value", []interface{}{'i', 'a', 'm', '4', 'p', 'q'}, '4'},
		{"float64 min value", []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7}, 0.5},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			bt := binarytree.NewBinaryTree()

			for i := 0; i < len(test.elements); i++ {
				err := bt.Insert(test.elements[i])
				if err != nil {
					t.Errorf("did not expect error %v", err)
				}
			}

			min := bt.GetMinValue()

			switch expMin := test.min.(type) {
			case bool:
				min, ok := min.(bool)
				if !ok {
					t.Error("expected bool type, did not get it")
				}

				if min != expMin {
					t.Errorf("expected %v, got %v", expMin, min)
				}
			case string:
				min, ok := min.(string)
				if !ok {
					t.Error("expected string type, did not get it")
				}

				if min != expMin {
					t.Errorf("expected %v, got %v", expMin, min)
				}
			case int:
				min, ok := min.(int)
				if !ok {
					t.Error("expected int type, did not get it")
				}

				if min != expMin {
					t.Errorf("expected %v, got %v", expMin, min)
				}
			case rune:
				min, ok := min.(rune)
				if !ok {
					t.Error("expected rune type, did not get it")
				}

				if min != expMin {
					t.Errorf("expected %v, got %v", expMin, min)
				}
			case float64:
				min, ok := min.(float64)
				if !ok {
					t.Error("expected float64 type, did not get it")
				}

				if min != expMin {
					t.Errorf("expected %v, got %v", expMin, min)
				}
			default:
				t.Error("unexpected min value type")
			}
		})
	}
}

func TestGetMaxValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
		max      interface{}
	}{
		{"bool max value", []interface{}{true, false}, true},
		{"string max value", []interface{}{"apple", "tea", "uno", "pi", "phone", "dino"}, "uno"},
		{"int max value", []interface{}{-54, 458, 23, -987, 1, 65, 87, 13, 990, 55, 7}, 990},
		{"rune max value", []interface{}{'i', 'a', 'm', '4', 'p', 'q'}, 'q'},
		{"float64 max value", []interface{}{5.5, 4.56789, 2.1, 0.5, 6.66, 987.7}, 987.7},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			bt := binarytree.NewBinaryTree()

			for i := 0; i < len(test.elements); i++ {
				err := bt.Insert(test.elements[i])
				if err != nil {
					t.Errorf("did not expect error %v", err)
				}
			}

			max := bt.GetMaxValue()

			switch expMax := test.max.(type) {
			case bool:
				max, ok := max.(bool)
				if !ok {
					t.Error("expected bool type, did not get it")
				}

				if max != expMax {
					t.Errorf("expected %v, got %v", expMax, max)
				}
			case string:
				max, ok := max.(string)
				if !ok {
					t.Error("expected string type, did not get it")
				}

				if max != expMax {
					t.Errorf("expected %v, got %v", expMax, max)
				}
			case int:
				max, ok := max.(int)
				if !ok {
					t.Error("expected int type, did not get it")
				}

				if max != expMax {
					t.Errorf("expected %v, got %v", expMax, max)
				}
			case rune:
				max, ok := max.(rune)
				if !ok {
					t.Errorf("expected rune type, did not get it")
				}

				if max != expMax {
					t.Errorf("expected %v, got %v", expMax, max)
				}
			case float64:
				max, ok := max.(float64)
				if !ok {
					t.Errorf("expected float64 type, did not get it")
				}

				if max != expMax {
					t.Errorf("expected %v, got %v", expMax, max)
				}
			default:
				t.Errorf("unexpected max value type")
			}
		})
	}
}

func TestInsertExistingValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		elements []interface{}
	}{
		{"bool existing value insert", []interface{}{true, true}},
		{"string existing value insert", []interface{}{"apple", "apple"}},
		{"int existing value insert", []interface{}{23, 23}},
		{"rune existing value insert", []interface{}{'m', 'm'}},
		{"float64 existing value insert", []interface{}{4.56789, 4.56789}},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.testName, func(t *testing.T) {
			t.Parallel()

			bt := binarytree.NewBinaryTree()

			err := bt.Insert(test.elements[0])
			if err != nil {
				t.Errorf("did not expect error %v", err)
			}

			err = bt.Insert(test.elements[1])
			if !errors.Is(err, binarytree.ErrNodeExists) {
				t.Error("expected ErrNodeExists, got none")
			}
		})
	}
}
