package unmarshaler_test

import (
	"encoding/json"
	"testing"

	"github.com/nndergunov/basicGo/JSON/custom/unmarshaler"
)

func TestMarshal(t *testing.T) {
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
			data: "russian military warship go _ yourself",
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
