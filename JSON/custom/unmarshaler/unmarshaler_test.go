package unmarshaler_test

import (
	"encoding/json"
	"testing"

	"github.com/nndergunov/basicGo/JSON/custom/unmarshaler"
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		data interface{}
	}{
		{
			name: "single int test",
			data: 5,
		},
		{
			name: "single float64 test",
			data: 3.4,
		},
		{
			name: "single bool test",
			data: true,
		},
		{
			name: "single string test",
			data: "russian military warship go yourself",
		},

		{
			name: "int slice test",
			data: []int{5, 4, 3, 2, 1, 0},
		},
		{
			name: "float64 slice test",
			data: []float64{5.4, 3.2, 0.1},
		},
		{
			name: "bool slice test",
			data: []bool{true, true, false, true},
		},
		{
			name: "string slice test",
			data: []string{"russian", "military", "warship", "go", "yourself"},
		},
		{
			name: "multiple type slice test",
			data: []interface{}{5, 4.3, true, "rmwgfy"},
		},
		{
			name: "slice in slice test",
			data: []interface{}{5, 4.3, true, []int{1, 2, 3}, "rmwgfy"},
		},
		{
			name: "map in slice test",
			data: []interface{}{5, 4.3, true, map[string]int{"one": 1, "two": 2, "three": 3}, "rmwgfy"},
		},
		{
			name: "map of int test",
			data: map[string]int{"one": 1, "two": 2, "three": 3},
		},
		{
			name: "map of float64 test",
			data: map[string]float64{"one.one": 1.1, "two.two": 2.2, "three.three": 3.3},
		},
		{
			name: "map of bool test",
			data: map[string]bool{"2+2=4": true, "Kyiv is the capital of Ukraine": true, "false": false},
		},
		{
			name: "map of strings test",
			data: map[string]string{"russian": "military", "warship": "go", "": "yourself"},
		},
		{
			name: "map of multiple types test",
			data: map[string]interface{}{"int": 1, "float64": 2.2, "bool": true, "string": "six"},
		},
		{
			name: "map of slices",
			data: map[string]interface{}{
				"one-three":  []int{1, 2, 3},
				"four-six":   []int{4, 5, 6},
				"seven-nine": []int{7, 8, 9},
			},
		},
		{
			name: "map of maps",
			data: map[string]interface{}{
				"one-three": map[string]int{
					"one":   1,
					"two":   2,
					"three": 3,
				},
				"four-six": map[string]int{
					"four": 4,
					"five": 5,
					"six":  6,
				},
				"seven-nine": map[string]int{
					"seven": 7,
					"eight": 8,
					"nine":  9,
				},
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(test.data)
			if err != nil {
				t.Error(err)
			}

			res, resType, err := unmarshaler.Unmarshal(data)
			if err != nil {
				t.Error(err)
			}

			switch exp := test.data.(type) {
			case int:
				if res.(float64) != float64(exp) || resType != "float64" {
					t.Errorf("expected: %v of type \"float64\", got: %v of type \"%s\"", exp, res, resType)
				}
			case float64:
				if res != test.data || resType != "float64" {
					t.Errorf("expected: %v of type \"float64\", got: %v of type \"%s\"", exp, res, resType)
				}
			case string:
				if res != test.data || resType != "string" {
					t.Errorf("expected: %v of type \"string\", got: %v of type \"%s\"", exp, res, resType)
				}
			case bool:
				if res != test.data || resType != "bool" {
					t.Errorf("expected: %v of type \"bool\", got: %v of type \"%s\"", exp, res, resType)
				}
			}
		})
	}
}
