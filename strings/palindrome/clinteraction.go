package main

import (
	"bufio"
	"fmt"
	"unicode"
)

// write puts line in output.
func newWriter(out *bufio.Writer) func(s string) error {
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

// read gets line from input and returns whole line and separated in two halves.
func newReader(in *bufio.Reader) func() ([]rune, error) {
	return func() ([]rune, error) {
		b, _, err := in.ReadLine()
		if err != nil {
			return nil, fmt.Errorf("read: %w", err)
		}

		word := []rune(string(b))

		if !isAWord(word) {
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