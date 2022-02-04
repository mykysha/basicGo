package main

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/nndergunov/basicGo/strings/anagram/anagramfinder"
)

func TestAnagrams(t *testing.T) {
	t.Parallel()

	table := struct {
		anagrams    [][]byte
		nonAnagrams [][]byte
	}{}

	f, err := os.Open("anagramexamples.txt")
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

		table.anagrams = append(table.anagrams, byteWord)

		table.nonAnagrams = append(table.nonAnagrams, byteWord[1:])
	}

	for i := 0; i < len(table.anagrams)-1; i++ {
		t.Run("palindrome_"+string(table.anagrams[i]), func(t *testing.T) {
			t.Parallel()

			words := strings.Split(string(table.anagrams[i]), " ")
			firstWord := []rune(words[0])
			secondWord := []rune(words[1])

			a := anagramfinder.NewAnagramFinder()
			res := a.Compute(firstWord, secondWord)

			if !res {
				t.Errorf("words \"%s\" and \"%s\" are anagrams, func returned false",
					string(firstWord), string(secondWord))
			}
		})

		t.Run("non-palindrome_"+string(table.nonAnagrams[i]), func(t *testing.T) {
			t.Parallel()

			words := strings.Split(string(table.nonAnagrams[i]), " ")
			firstWord := []rune(words[0])
			secondWord := []rune(words[1])

			a := anagramfinder.NewAnagramFinder()
			res := a.Compute(firstWord, secondWord)

			if res {
				t.Errorf("words \"%s\" and \"%s\" are not anagrams, func returned true",
					string(firstWord), string(secondWord))
			}
		})
	}
}
