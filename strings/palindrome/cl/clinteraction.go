package cl

import (
	"bufio"
	"fmt"
)

// NewWriter returns write func.
func NewWriter(out *bufio.Writer) func(s string) error {
	return func(s string) error {
		_, err := out.WriteString(s)
		if err != nil {
			return fmt.Errorf("writing: %w", err)
		}

		err = out.Flush()
		if err != nil {
			return fmt.Errorf("flushing: %w", err)
		}

		return nil
	}
}

// NewReader returns read func.
func NewReader(in *bufio.Reader) func() (string, error) {
	return func() (string, error) {
		b, _, err := in.ReadLine()
		if err != nil {
			return "", fmt.Errorf("read: %w", err)
		}

		word := string(b)

		return word, nil
	}
}
