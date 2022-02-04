package main

import (
	"bufio"
	"github.com/nndergunov/basicGo/strings/palindrome/pool"
	"io"
	"os"
	"testing"
)

func TestPalindromesSingleThread(t *testing.T) {
	t.Parallel()

	table := struct {
		palindromes    [][]rune
		nonPalindromes [][]rune
	}{}

	f, err := os.Open("palindromeexamples.txt")
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(f)

	for {
		byteWord, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			t.Fatal(err)
		}

		word := []rune(string(byteWord))

		table.palindromes = append(table.palindromes, word)

		table.nonPalindromes = append(table.nonPalindromes, word[1:])
	}

	for i := 0; i < len(table.palindromes)-1; i++ {
		t.Run("palindrome_"+string(table.palindromes[i]), func(t *testing.T) {
			t.Parallel()

			word := table.palindromes[i]

			res := checkIfPalindrome(word)

			if !res {
				t.Errorf("word \"%s\" is a palindrome, func returned false", string(word))
			}
		})

		t.Run("non-palindrome_"+string(table.nonPalindromes[i]), func(t *testing.T) {
			t.Parallel()

			word := table.nonPalindromes[i]

			res := checkIfPalindrome(word)

			if res {
				t.Errorf("word \"%s\" is a- not palindrome, func returned true", string(word))
			}
		})
	}
}

func TestPalindromesMultiThread(t *testing.T) {
	t.Parallel()

	table := struct {
		workerNumber    int
		substringLength int
		palindromes     [][]rune
		nonPalindromes  [][]rune
	}{}

	table.workerNumber = 4
	table.substringLength = 5

	f, err := os.Open("palindromeexamples.txt")
	if err != nil {
		t.Fatal(err)
	}

	reader := bufio.NewReader(f)

	for {
		byteWord, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			t.Fatal(err)
		}

		word := []rune(string(byteWord))

		table.palindromes = append(table.palindromes, word)

		table.nonPalindromes = append(table.nonPalindromes, word[1:])
	}

	for i := 0; i < len(table.palindromes)-1; i++ {
		t.Run("palindrome_"+string(table.palindromes[i]), func(t *testing.T) {
			t.Parallel()

			word := table.palindromes[i]

			p := pool.NewPool(table.workerNumber)

			p.Start()

			s := makeSubstrings(word, table.substringLength)

			resultChan := make(chan bool, len(s))

			for _, val := range s {
				substring := val

				task := func() {
					resultChan <- checkIfPalindrome(substring)
				}

				p.AddTask(task)
			}

			for i := 0; i < len(s); i++ {
				res := <-resultChan

				if !res {
					t.Errorf("substring \"%s\" is a palindrome, func returned false", string(word))
				}
			}
		})

		t.Run("non-palindrome_"+string(table.nonPalindromes[i]), func(t *testing.T) {
			t.Parallel()

			word := table.nonPalindromes[i]

			p := pool.NewPool(table.workerNumber)

			p.Start()

			s := makeSubstrings(word, table.substringLength)

			resultChan := make(chan bool, len(s))

			for _, val := range s {
				substring := val

				task := func() {
					resultChan <- checkIfPalindrome(substring)
				}

				p.AddTask(task)
			}

			result := true

			for i := 0; i < len(s); i++ {
				res := <-resultChan

				if !res {
					result = false
				}
			}

			if result {
				t.Errorf("substring \"%s\" is not a palindrome, func returned true", string(word))
			}
		})
	}
}