package main

import (
	"testing"
)

// given the lack of support for multiple-typed struct array in Go,
// we have to use old method of interfaces and type assertions
func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		data     interface{}
		expected interface{}
	}{
		{
			name:     "bool slice",
			data:     []bool{true, true, true, false, false},
			expected: []bool{true, false},
		},
		{
			name:     "int slice",
			data:     []int{1, 1, 2, 2, 3, 3, 2, 4, 4, 3, 5, 5, 4, 3, 4, 4, 5, 5, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "string slice",
			data:     []string{"one", "two", "two", "three", "three", "three"},
			expected: []string{"one", "two", "three"},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			switch arr := test.data.(type) {
			case []bool:
				exp := test.expected.([]bool)

				res := RemoveDuplicates(arr)

				for i := 0; i < len(exp); i++ {
					if res[i] != exp[i] {
						t.Errorf("expected: %v, got: %v", test.expected, res)

						break
					}
				}
			case []int:
				exp := test.expected.([]int)

				res := RemoveDuplicates(arr)

				for i := 0; i < len(exp); i++ {
					if res[i] != exp[i] {
						t.Errorf("expected: %v, got: %v", test.expected, res)

						break
					}
				}
			case []string:
				exp := test.expected.([]string)

				res := RemoveDuplicates(arr)

				for i := 0; i < len(exp); i++ {
					if res[i] != exp[i] {
						t.Errorf("expected: %v, got: %v", test.expected, res)

						break
					}
				}
			default:
				t.Error("test data is not of comparable type")
			}
		})
	}
}
