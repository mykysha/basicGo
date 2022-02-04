package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/nndergunov/basicGo/strings/palindrome/pool"
	"log"
	"os"
	"unicode"
)

var errNotAWord = errors.New("the string that was read is not a word")

func main() {
	out := bufio.NewWriter(os.Stdout)
	writer := newWriter(out)
	in := bufio.NewReader(os.Stdin)
	reader := newReader(in)

	msg := fmt.Sprint("This program will tell if your word is a palindrome\n" +
		"Palindrome is a word that reads equally from each side\n" +
		"Please enter your word:\t")

	err := writer(msg)
	if err != nil {
		log.Println(err)
	}

	word, err := reader()
	if err != nil {
		log.Println(err)
	}

	workerNumber := 4
	p := pool.NewPool(workerNumber)

	p.Start()

	substringLength := 20
	substrings := makeSubstrings(word, substringLength)

	resultChan := make(chan bool, len(substrings))

	for _, val := range substrings {
		subword := val

		task := func() {
			resultChan <- checkIfPalindrome(subword)
		}

		p.AddTask(task)
	}

	res := fmt.Sprintf("The word \"%s\" is", string(word))

	for result := range resultChan {
		if !result {
			res += " not"
		}
	}

	res += " a palindrome"

	err = writer(res)
	if err != nil {
		log.Println(err)
	}
}

// makeSubstrings divides big words into smaller ones combining parts that are located on both sides of the word.
// substringLength will be round up to the closest even number.
// Example: makeSubstrings("aaabbbcccddd", 6) will return aaaddd, bbbccc.
func makeSubstrings(word []rune, substringLength int) [][]rune {
	var substrings [][]rune

	if substringLength < 2 {
		substringLength = 2
	}

	substringLength /= 2

	firstHalf := word[:len(word)/2]
	secondHalf := word[len(word)/2:]

	for from, till := 0, substringLength; till < len(firstHalf); from, till = from+substringLength, till+substringLength {
		substring := make([]rune, substringLength*2)

		firstSubstring := firstHalf[from:till]
		secondSubstring := secondHalf[len(secondHalf)-till : len(secondHalf)-from]

		copy(substring, firstSubstring)
		copy(substring[10:], secondSubstring)

		substrings = append(substrings, substring)
	}

	tail := len(firstHalf) % substringLength
	wrote := len(firstHalf) - tail

	lastSubstring := make([]rune, tail*2)

	copy(lastSubstring, firstHalf[wrote:])
	copy(lastSubstring[tail:], secondHalf[:len(secondHalf)-wrote])

	substrings = append(substrings, lastSubstring)

	return substrings
}

// checkIfPalindrome finds out can the given []rune be read backwards or not.
func checkIfPalindrome(word []rune) bool {
	for i := 0; i < len(word)/2; i++ {
		r1 := unicode.ToLower(word[i])
		r2 := unicode.ToLower(word[len(word)-i-1])

		if r1 != r2 {
			return false
		}
	}

	return true
}