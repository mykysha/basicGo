package anagramfinder

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

var errNotAWord = errors.New("the string that was read is not a word")

func convertToRune(wordStr string) ([]rune, error) {
	var (
		word  []rune
		total int
	)

	wordLen := len(wordStr)

	for {
		symbol, runeCount := utf8.DecodeRuneInString(wordStr)

		if symbol == utf8.RuneError {
			if runeCount == 0 {
				break
			}

			return nil, fmt.Errorf("string to rune conversion")
		}

		word = append(word, symbol)

		total += runeCount

		if total == wordLen {
			break
		}

		wordStr = wordStr[runeCount:]
	}

	return word, nil
}

// isAWord checks if given rune array contains symbols that are not letters.
func isAWord(r []rune) bool {
	for i := 0; i < len(r); i++ {
		if !unicode.IsLetter(r[i]) {
			return false
		}
	}

	return true
}
