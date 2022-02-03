package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	s := sortingComparison{}

	s.Init(out)

	s.Start()
}

type sortingComparison struct {
	out                *bufio.Writer
	builtInSortTime    time.Duration
	bubbleSortTime     time.Duration
	randomJokeSortTime time.Duration
}

func (s *sortingComparison) Init(out *bufio.Writer) {
	s.out = out
}

func (s *sortingComparison) Start() {
	msg := fmt.Sprintf("Hello! This program will measure the time it takes to sort\n" +
		"slice which contains integers from 10000 to 1\n" +
		"using different algorithms.")

	s.writeLine(msg)

	numberOfAlgorithms := 3
	wg := new(sync.WaitGroup)

	wg.Add(numberOfAlgorithms)

	go func(wg *sync.WaitGroup) {
		s.builtInSort()

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		s.bubbleSort()

		wg.Done()
	}(wg)

	go func(wg *sync.WaitGroup) {
		s.randomJokeSort()

		wg.Done()
	}(wg)

	wg.Wait()

	s.writeResults()
}

func (s *sortingComparison) builtInSort() {
	testSlice := generateSlice()

	start := time.Now()

	sort.Ints(testSlice)

	s.builtInSortTime = time.Since(start)
}

func (s *sortingComparison) bubbleSort() {
	testSlice := generateSlice()

	start := time.Now()
	for i := len(testSlice) - 1; i > 0; i-- {
		swapped := false

		for j := 0; j < i; j++ {
			if testSlice[j] > testSlice[j+1] {
				testSlice[j], testSlice[j+1] = testSlice[j+1], testSlice[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	s.bubbleSortTime = time.Since(start)
}

func (s *sortingComparison) randomJokeSort() {
	testSlice := generateSlice()
	secondLimit := time.Duration(5)

	start := time.Now()

	for {
		rand.Seed(time.Now().UnixNano())

		rand.Shuffle(len(testSlice), func(i, j int) { testSlice[i], testSlice[j] = testSlice[j], testSlice[i] })

		success := true

		for i := 1; i < len(testSlice); i++ {
			if testSlice[i] < testSlice[i-1] {
				success = false
				break
			}
		}

		if success == true {
			s.randomJokeSortTime = time.Since(start)
			break
		}

		if time.Since(start) > secondLimit*time.Second {
			failTime := time.Duration(999)

			s.randomJokeSortTime = failTime * time.Hour
		}
	}
}

func (s sortingComparison) writeResults() {
	builtIn := fmt.Sprintf("built in sort package time: %v", s.builtInSortTime)

	s.writeLine(builtIn)

	bubble := fmt.Sprintf("bubble sort time: %v", s.bubbleSortTime)

	s.writeLine(bubble)

	random := fmt.Sprintf("random shuffle of slice sort time: %v", s.randomJokeSortTime)

	s.writeLine(random)
}

func (s sortingComparison) writeLine(msg string) {
	for {
		_, err := s.out.WriteString(msg + "\n")
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}

	for {
		err := s.out.Flush()
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
}

func generateSlice() []int {
	length := 100000

	s := make([]int, 0, length)

	for i := length; i > 0; i-- {
		s = append(s, i)
	}

	return s
}