package palindromefinder

import (
	"fmt"

	"github.com/nndergunov/basicGo/strings/palindrome/pool"
)

func CheckIfPalindrome(wordStr string) (bool, error) {
	word, err := convertToRune(wordStr)
	if err != nil {
		return false, fmt.Errorf("string to rune conversion: %w", err)
	}

	if !isAWord(word) {
		return false, fmt.Errorf("symbol check: %w", errNotAWord)
	}

	return checkIfPalindrome(word), nil
}

func CheckIfPalindromeConcurrently(wordStr string) (bool, error) {
	word, err := convertToRune(wordStr)
	if err != nil {
		return false, fmt.Errorf("string to rune conversion: %w", err)
	}

	workerNumber := 4
	p := pool.NewPool(workerNumber)

	p.Start()

	substringLength := 10
	substrings := makeSubstrings(word, substringLength)

	resultChan := make(chan bool, len(substrings))

	for _, val := range substrings {
		subword := val

		task := func() {
			if !isAWord(subword) {
				resultChan <- false

				return
			}

			resultChan <- checkIfPalindrome(subword)
		}

		p.AddTask(task)
	}

	for i := 0; i < len(substrings); i++ {
		result := <-resultChan

		if !result {
			p.Stop()

			return false, nil
		}
	}

	return true, nil
}

// makeSubstrings divides big words into smaller ones combining parts that are located on both sides of the word.
// substringLength will be round up to the closest even number.
// Example: makeSubstrings("aaabbbcccddd", 6) will return aaaddd, bbbccc.
func makeSubstrings(word []rune, substringLength int) [][]rune {
	var substrings [][]rune

	if len(word) <= substringLength {
		substrings = append(substrings, word)

		return substrings
	}

	if minimalSubstringLength := 2; substringLength < minimalSubstringLength {
		substringLength = minimalSubstringLength
	}

	numberOfHalves := 2

	substringLength /= numberOfHalves

	firstHalf := word[:len(word)/numberOfHalves]
	secondHalf := word[(len(word) - len(word)/2):]

	for from, till := 0, substringLength; till <= len(firstHalf); from, till = from+substringLength, till+substringLength {
		substring := make([]rune, substringLength*numberOfHalves)

		firstSubstring := firstHalf[from:till]
		secondSubstring := secondHalf[len(secondHalf)-till : len(secondHalf)-from]

		copy(substring, firstSubstring)
		copy(substring[substringLength:], secondSubstring)

		substrings = append(substrings, substring)
	}

	if len(word)%substringLength != 0 {
		tail := len(firstHalf) % substringLength
		wrote := len(firstHalf) - tail

		lastSubstring := make([]rune, tail*numberOfHalves)

		copy(lastSubstring, firstHalf[wrote:])
		copy(lastSubstring[tail:], secondHalf[:len(secondHalf)-wrote])

		substrings = append(substrings, lastSubstring)
	}

	return substrings
}
