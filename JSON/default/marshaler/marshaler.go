package marshaler

import (
	"encoding/json"
	"fmt"
)

func DefaultMarshal(req interface{}) ([]byte, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("default marshal: %w", err)
	}

	return data, nil
}
