package unmarshaler

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrNotJSON = errors.New("given data slice is not a valid JSON struct and can not be regarded so")

func DefaultUnmarshal(data []byte, res interface{}) error {
	if !json.Valid(data) {
		return ErrNotJSON
	}

	err := json.Unmarshal(data, res)
	if err != nil {
		return fmt.Errorf("default unmarshal: %w", err)
	}

	return nil
}

func IntUnmarshal(data []byte) (int, error) {
	if !json.Valid(data) {
		return 0, ErrNotJSON
	}

	var res int

	err := json.Unmarshal(data, &res)
	if err != nil {
		return 0, fmt.Errorf("default int unmarshal: %w", err)
	}

	return res, nil
}

func BoolUnmarshal(data []byte) (bool, error) {
	if !json.Valid(data) {
		return false, ErrNotJSON
	}

	var res bool

	err := json.Unmarshal(data, &res)
	if err != nil {
		return false, fmt.Errorf("default bool unmarshal: %w", err)
	}

	return res, nil
}

func StringUnmarshal(data []byte) (string, error) {
	if !json.Valid(data) {
		return "", ErrNotJSON
	}

	var res string

	err := json.Unmarshal(data, &res)
	if err != nil {
		return "", fmt.Errorf("default string unmarshal: %w", err)
	}

	return res, nil
}
