package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
	"unicode"
)

/*
	Purely algorithmic task. Have to find the way to efficiently sort given strings, may be make it concurrent(?)
	May be possible to read the whole thing but other symbols are not accounted for.
	Will be useful to count the sum of runes/bytes to find if it is even remotely possible for them to be equal.
*/

var errNotAWord = errors.New("the string that was read is not a word")

func main() {
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	a := anagramFinder{}
	a.Init(in, out)
}

type anagramFinder struct {
	in                   *bufio.Reader
	out                  *bufio.Writer
	firstWord            []rune
	secondWord           []rune
	firstWordByLetters   map[rune]uint
	secondWordByLetters  map[rune]uint
	firstWordByteLength  int
	secondWordByteLength int
}

func (a *anagramFinder) Init(in *bufio.Reader, out *bufio.Writer) {
	a.in = in
	a.out = out
	a.firstWordByLetters = make(map[rune]uint)
	a.secondWordByLetters = make(map[rune]uint)

	a.start()
}

func (a *anagramFinder) start() {
	msg := fmt.Sprint("This program will tell if your words are an anagram.\n" +
		"Anagrams are words that are made from the same letters.\n")

	err := a.write(msg)
	if err != nil {
		log.Println(err)
	}

	shouldCheckFurther := a.getWordInfo()

	if !shouldCheckFurther {
		a.negativeResult()

		return
	}

	wg := new(sync.WaitGroup)
	numberOfWords := 2

	wg.Add(numberOfWords)

	go func(wg *sync.WaitGroup) {
		a.firstWordByLetters = sort(a.firstWord)

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		a.secondWordByLetters = sort(a.secondWord)

		wg.Done()
	}(wg)

	wg.Wait()

	if len(a.firstWordByLetters) != len(a.secondWordByLetters) {
		a.negativeResult()

		return
	}

	for key, val := range a.firstWordByLetters {
		if val != a.secondWordByLetters[key] {
			a.negativeResult()

			return
		}
	}

	a.positiveResult()
}

func (a *anagramFinder) getWordInfo() bool {
	a.getFirstWord()

	a.getSecondWord()

	if len(a.firstWord) != len(a.secondWord) {
		return false
	}

	wg := new(sync.WaitGroup)
	numberOfWords := 2

	wg.Add(numberOfWords)

	go func(wg *sync.WaitGroup) {
		a.firstWordByteLength = findByteLength(a.firstWord)

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		a.secondWordByteLength = findByteLength(a.secondWord)

		wg.Done()
	}(wg)

	wg.Wait()

	if a.firstWordByteLength != a.secondWordByteLength {
		return false
	}

	return true
}

func sort(r []rune) map[rune]uint {
	letterMap := make(map[rune]uint)

	for _, val := range r {
		letterMap[val]++
	}

	return letterMap
}

func (a *anagramFinder) getFirstWord() {
	msg := "Please enter your first word:\t"

	for {
		err := a.write(msg)
		if err != nil {
			log.Println(err)

			continue
		}

		a.firstWord, err = a.read()
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
}

func (a *anagramFinder) getSecondWord() {
	msg := "Please enter your second word:\t"
	for {
		err := a.write(msg)
		if err != nil {
			log.Println(err)

			continue
		}

		a.secondWord, err = a.read()
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
}

// write puts line in output.
func (a anagramFinder) write(s string) error {
	_, err := a.out.WriteString(s)
	if err != nil {
		return fmt.Errorf("writing: %w", err)
	}

	err = a.out.Flush()
	if err != nil {
		return fmt.Errorf("flushing: %w", err)
	}

	return nil
}

// read gets line from input.
func (a anagramFinder) read() ([]rune, error) {
	b, _, err := a.in.ReadLine()
	if err != nil {
		return nil, fmt.Errorf("read: %w", err)
	}

	word := []rune(string(b))

	if isAWord(word) == false {
		return nil, fmt.Errorf("examine: %w", errNotAWord)
	}

	return word, nil
}

func (a anagramFinder) negativeResult() {
	msg := fmt.Sprintf("Words \"%s\" and \"%s\" are not anagrams.", string(a.firstWord), string(a.secondWord))

	for {
		err := a.write(msg)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
}

func (a anagramFinder) positiveResult() {
	msg := fmt.Sprintf("Words \"%s\" and \"%s\" are anagrams.", string(a.firstWord), string(a.secondWord))

	for {
		err := a.write(msg)
		if err != nil {
			log.Println(err)
		} else {
			break
		}
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

func findByteLength(r []rune) int {
	var res int

	for _, val := range r {
		res += int(val)
	}

	return res
}