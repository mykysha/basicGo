package marshaler_test

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/nndergunov/basicGo/JSON/default/marshaler"
)

func TestDefaultMarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		request interface{}
	}{
		{
			name:    "int default marshal",
			request: 2,
		},
		{
			name:    "bool default marshal",
			request: true,
		},
		{
			name:    "string default marshal",
			request: "russian military warship go _ yourself",
		},
		{
			name:    "slice default marshal",
			request: interface{}([]string{"russian", "military", "warship", "go", "", "yourself"}),
		},
		{
			name:    "map default marshal",
			request: map[int]string{0: "russian", 1: "military", 2: "warship", 3: "go", 4: "", 5: "yourself"},
		},
		{
			name: "struct without tags (all fields uppercase) default marshal",
			request: struct {
				Name   string
				Age    int
				Gender string
			}{
				Name:   "John Smith",
				Age:    21,
				Gender: "male",
			},
		},
		{
			name: "struct without tags (some fields lowercase) default marshal",
			request: struct {
				Name   string
				Age    int
				gender string
			}{
				Name:   "John Smith",
				Age:    21,
				gender: "male",
			},
		},
		{
			name: "struct without tags (all fields lowercase) default marshal",
			request: struct {
				name   string
				age    int
				gender string
			}{
				name:   "John Smith",
				age:    21,
				gender: "male",
			},
		},
		{
			name: "struct with all tags (all fields uppercase) default marshal",
			request: struct {
				Name   string `json:"fullName"`
				Age    int    `json:"givenAge"`
				Gender string `json:"chosenGender"`
			}{
				Name:   "John Smith",
				Age:    21,
				Gender: "male",
			},
		},
		{
			name: "struct with all tags (some fields lowercase) default marshal",
			request: struct {
				Name   string `json:"fullName"`
				Age    int    `json:"givenAge"`
				gender string `json:"chosenGender"` /* go built-in check is not happy with this implementation
				+ govet linter triggered */
			}{
				Name:   "John Smith",
				Age:    21,
				gender: "male",
			},
		},
		{
			name: "struct with some tags (all fields uppercase) default marshal",
			request: struct {
				Name   string `json:"fullName"`
				Age    int    `json:"givenAge"`
				Gender string
			}{
				Name:   "John Smith",
				Age:    21,
				Gender: "male",
			},
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			data, err := marshaler.DefaultMarshal(test.request)
			if err != nil {
				t.Errorf("expected no error, got: %v", err)
			}

			log.Printf("%s: %s", test.name, data)

			if !json.Valid(data) {
				t.Errorf("json marshal failed")
			}
		})
	}
}
