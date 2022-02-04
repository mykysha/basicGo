package cl

import (
	"bufio"
	"errors"
	"fmt"
	"unicode"
)

var errNotAWord = errors.New("the string that was read is not a word")

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
func NewReader(in *bufio.Reader) func() ([]rune, error) {
	return func() ([]rune, error) {
		b, _, err := in.ReadLine()
		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		word := []rune(string(b))

		if isAWord(word) == false {
			return nil, fmt.Errorf("examine: %w", errNotAWord)
		}

		return word, nil
	}
}

// isAWord checks if given rune array contains symbols that are not letters. Might use workerpool as well to speed up.
func isAWord(r []rune) bool {
	for i := 0; i < len(r); i++ {
		if !unicode.IsLetter(r[i]) {
			return false
		}
	}

	return true
}
