package marshaler

import (
	"fmt"
	"strings"
)

const supportedTypes = "int, float64, bool, rune, string, slices of those types, map[string]"

var ErrUnsupportedType = fmt.Errorf("json marshal unexpected type. supported types: %s", supportedTypes)

// Marshal is used to convert int, float64, bool, rune, string, slices of those types, map[string] to JSON.
//
// Possible usage of composite higher-level types (slice of slices/map of maps) using interface{} type.
func Marshal(object interface{}) (data []byte, err error) {
	switch object.(type) {
	case map[string]int, map[string]float64, map[string]bool, map[string]rune, map[string]string, map[string]interface{}:
		mapData, err := mapMarshal(object)
		if err != nil {
			return nil, fmt.Errorf("map marshal: %w", err)
		}

		return mapData, nil
	case []int, []float64, []bool, []rune, []string, []interface{}: // todo reflect package.
		arrayData, err := arrayMarshal(object)
		if err != nil {
			return nil, fmt.Errorf("array marshal: %w", err)
		}

		return arrayData, nil
	case int, float64, bool, rune, string:
		basicData := basicMarshal(object)

		return basicData, nil
	default:
		return nil, ErrUnsupportedType
	}
}

func mapMarshal(stringMap interface{}) ([]byte, error) {
	switch m := stringMap.(type) {
	case map[string]interface{}:
		jsonObject := "{"

		for key, val := range m {
			jsonObject += fmt.Sprintf("\"%s\": ", key)

			el, err := Marshal(val)
			if err != nil {
				return nil, fmt.Errorf("element marshal: %w", err)
			}

			jsonObject += fmt.Sprintf("%s, ", string(el))
		}

		jsonObject = strings.TrimSuffix(jsonObject, ", ")

		jsonObject += "}"

		return []byte(jsonObject), nil
	default:
		interfaceMap := convertToInterfaceMap(stringMap)

		return mapMarshal(interfaceMap)
	}
}

func arrayMarshal(array interface{}) ([]byte, error) {
	switch el := array.(type) {
	case []interface{}:
		jsonArray := "["

		for i := 0; i < len(el); i++ {
			jsonElement, err := Marshal(el[i])
			if err != nil {
				return nil, fmt.Errorf("element marshal: %w", err)
			}

			jsonArray += string(jsonElement)

			if i+1 != len(el) {
				jsonArray += ", "
			}
		}

		jsonArray += "]"

		return []byte(jsonArray), nil
	default:
		interfaceSlice := convertToInterfaceSlice(array)

		return arrayMarshal(interfaceSlice)
	}
}

func basicMarshal(element interface{}) []byte {
	switch el := element.(type) {
	case int, float64, bool, rune:
		return []byte(fmt.Sprintf("%v", el))
	case string:
		return []byte(fmt.Sprintf("\"%s\"", el))
	default:
		return nil
	}
}

func convertToInterfaceMap(i interface{}) map[string]interface{} {
	interfaceMap := make(map[string]interface{})

	switch givenMap := i.(type) {
	case map[string]int:
		for key, el := range givenMap {
			interfaceMap[key] = el
		}

		return interfaceMap
	case map[string]float64:
		for key, el := range givenMap {
			interfaceMap[key] = el
		}

		return interfaceMap
	case map[string]bool:
		for key, el := range givenMap {
			interfaceMap[key] = el
		}

		return interfaceMap
	case map[string]rune:
		for key, el := range givenMap {
			interfaceMap[key] = el
		}

		return interfaceMap
	case map[string]string:
		for key, el := range givenMap {
			interfaceMap[key] = el
		}

		return interfaceMap
	default:
		return nil
	}
}

func convertToInterfaceSlice(i interface{}) []interface{} {
	switch arr := i.(type) {
	case []int:
		s := make([]interface{}, len(arr))
		for key, el := range arr {
			s[key] = el
		}

		return s
	case []float64:
		s := make([]interface{}, len(arr))
		for key, el := range arr {
			s[key] = el
		}

		return s
	case []bool:
		s := make([]interface{}, len(arr))
		for key, el := range arr {
			s[key] = el
		}

		return s
	case []rune:
		s := make([]interface{}, len(arr))
		for key, el := range arr {
			s[key] = el
		}

		return s
	case []string:
		s := make([]interface{}, len(arr))
		for key, el := range arr {
			s[key] = el
		}

		return s
	default:
		return nil
	}
}
