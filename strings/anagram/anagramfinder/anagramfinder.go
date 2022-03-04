package anagramfinder

import (
	"fmt"
	"sync"
)

func CheckIfAnagram(firstWordStr, secondWordStr string) (bool, error) {
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

	shouldCheckFurther := compareWordInfo(firstWord, secondWord)

	if !shouldCheckFurther {
		return false, nil
	}

	wg := new(sync.WaitGroup)
	numberOfWords := 2

	wg.Add(numberOfWords)

	lettersChan := make(chan map[rune]uint, numberOfWords)

	go func(wg *sync.WaitGroup) {
		lettersChan <- sort(firstWord)

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		lettersChan <- sort(secondWord)

		wg.Done()
	}(wg)

	wg.Wait()

	firstLetterMap := <-lettersChan
	secondLetterMap := <-lettersChan

	if len(firstLetterMap) != len(secondLetterMap) {
		return false, nil
	}

	for key, val := range firstLetterMap {
		if val != secondLetterMap[key] {
			return false, nil
		}
	}

	return true, nil
}

func compareWordInfo(firstWord, secondWord []rune) bool {
	if len(firstWord) != len(secondWord) {
		return false
	}

	wg := new(sync.WaitGroup)
	numberOfWords := 2

	wg.Add(numberOfWords)

	lengthChan := make(chan int, numberOfWords)

	go func(wg *sync.WaitGroup) {
		lengthChan <- findByteLength(firstWord)

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		lengthChan <- findByteLength(secondWord)

		wg.Done()
	}(wg)

	wg.Wait()

	firstByteLength := <-lengthChan
	secondByteLength := <-lengthChan

	return firstByteLength == secondByteLength
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
