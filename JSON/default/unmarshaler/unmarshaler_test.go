package unmarshaler_test

import (
	"encoding/json"
	"testing"

	"github.com/nndergunov/basicGo/JSON/default/unmarshaler"
)

func TestDefaultUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		request interface{}
	}{
		{
			name:    "int default unmarshal",
			request: 2,
		},
		{
			name:    "bool default unmarshal",
			request: true,
		},
		{
			name:    "string default unmarshal",
			request: "russian military warship go _ yourself",
		},
		{
			name:    "slice default unmarshal",
			request: interface{}([]string{"russian", "military", "warship", "go", "", "yourself"}),
		},
		{
			name:    "map default unmarshal",
			request: map[int]string{0: "russian", 1: "military", 2: "warship", 3: "go", 4: "", 5: "yourself"},
		},
		{
			name: "struct without tags, all fields uppercase default unmarshal",
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
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(test.request)
			if err != nil {
				t.Errorf("expected no error while marshaling, got: %v", err)
			}

			switch expectedData := test.request.(type) {
			case int:
				var res int

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				if res != expectedData {
					t.Errorf("expected: %v, got: %v", test.request, res)
				}
			case bool:
				var res bool

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				if res != expectedData {
					t.Errorf("expected: %v, got: %v", test.request, res)
				}
			case string:
				var res string

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				if res != expectedData {
					t.Errorf("expected: %v, got: %v", test.request, res)
				}
			case struct {
				Name   string
				Age    int
				Gender string
			}:
				var res struct {
					Name   string
					Age    int
					Gender string
				}

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				if res != test.request {
					t.Errorf("expected: %v, got: %v", test.request, res)
				}
			case map[int]string:
				res := make(map[int]string)

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				for key, val := range expectedData {
					if val != res[key] {
						t.Errorf("expected: %v, got: %v", test.request, res)
					}
				}
			case []string:
				res := make([]string, len(expectedData))

				err = unmarshaler.DefaultUnmarshal(data, &res)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}

				for key, val := range expectedData {
					if val != res[key] {
						t.Errorf("expected: %v, got: %v", test.request, res)
					}
				}
			}
		})
	}
}

func TestSpecifiedUnmarshal(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		request interface{}
	}{
		{
			name:    "int specified unmarshal",
			request: 2,
		},
		{
			name:    "bool specified unmarshal",
			request: true,
		},
		{
			name:    "string specified unmarshal",
			request: "russian military warship go _ yourself",
		},
	}

	for _, currentTest := range tests {
		test := currentTest

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(test.request)
			if err != nil {
				t.Errorf("expected no error while marshaling, got: %v", err)
			}

			var res interface{}

			switch test.request.(type) {
			case int:
				res, err = unmarshaler.IntUnmarshal(data)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}
			case bool:
				res, err = unmarshaler.BoolUnmarshal(data)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}
			case string:
				res, err = unmarshaler.StringUnmarshal(data)
				if err != nil {
					t.Errorf("expected no error while unmarshaling, got: %v", err)
				}
			}

			if res != test.request {
				t.Errorf("expected: %v, got: %v", test.request, res)
			}
		})
	}
}
