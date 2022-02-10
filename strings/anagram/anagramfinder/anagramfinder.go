package anagramfinder

import (
	"fmt"
	"sync"
	"time"
)

type AnagramFinder struct {
	firstWord            []rune
	secondWord           []rune
	firstWordByLetters   map[rune]uint
	secondWordByLetters  map[rune]uint
	firstWordByteLength  int
	secondWordByteLength int
	computing            bool
}

func NewAnagramFinder() *AnagramFinder {
	a := AnagramFinder{
		firstWord:            nil,
		secondWord:           nil,
		firstWordByLetters:   nil,
		secondWordByLetters:  nil,
		firstWordByteLength:  0,
		secondWordByteLength: 0,
		computing:            false,
	}

	a.firstWordByLetters = make(map[rune]uint)
	a.secondWordByLetters = make(map[rune]uint)

	return &a
}

func (a *AnagramFinder) Compute(firstWordStr, secondWordStr string) (bool, error) {
	firstWord, err := convertToRune(firstWordStr)
	if err != nil {
		return false, fmt.Errorf("first word to rune conversion: %w", err)
	}

	secondWord, err := convertToRune(secondWordStr)
	if err != nil {
		return false, fmt.Errorf("second word to rune conversion: %w", err)
	}

	if !isAWord(firstWord) || !isAWord(secondWord) {
		return false, fmt.Errorf("word analysis: %w", errNotAWord)
	}

	for {
		if a.computing {
			time.Sleep(1 * time.Second)

			continue
		}

		a.computing = true

		break
	}

	defer func() {
		a.computing = false
	}()

	a.firstWord = firstWord
	a.secondWord = secondWord

	shouldCheckFurther := a.compareWordInfo()

	if !shouldCheckFurther {
		return false, nil
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
		return false, nil
	}

	for key, val := range a.firstWordByLetters {
		if val != a.secondWordByLetters[key] {
			return false, nil
		}
	}

	return true, nil
}

func (a AnagramFinder) compareWordInfo() bool {
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

func findByteLength(r []rune) int {
	var res int

	for _, val := range r {
		res += int(val)
	}

	return res
}
