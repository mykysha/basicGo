package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"unicode"
)

var errNotAWord = errors.New("the string that was read is not a word")

func main() {
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	msg := fmt.Sprint("This program will tell if your word is a palindrome\n" +
		"Palindrome is a word that reads equally from each side\n" +
		"Please enter your word:\t")

	err := write(msg, out)
	if err != nil {
		log.Println(err)
	}

	word, firstHalf, secondHalf, err := read(in)
	if err != nil {
		log.Println(err)
	}

	res := fmt.Sprintf("The word \"%s\" is", word)

	if !isAPalindrome(firstHalf, secondHalf) {
		res += " not"
	}

	res += " a palindrome"

	err = write(res, out)
	if err != nil {
		log.Println(err)
	}
}

func isAPalindrome(firstHalf, secondHalf []rune) bool {
	workerNumber := 4

	substringMaxLength := 10
	substrings := makeSubstings(firstHalf, secondHalf, substringMaxLength)

	p := pool{
		workerNumber: workerNumber,
		workers:      nil,
		took:         0,
		substrings:   substrings,
		result:       true,
	}

	return p.poolCall()
}

func makeSubstings(firstHalf, secondHalf []rune, maxLength int) [][]rune {
	var substrings [][]rune

	for from, till := 0, maxLength; till < len(firstHalf); from, till = from+maxLength, till+maxLength {
		substring := make([]rune, maxLength*2)

		firstSubstring := firstHalf[from:till]
		secondSubstring := secondHalf[len(secondHalf)-till : len(secondHalf)-from]

		copy(substring, firstSubstring)
		copy(substring[10:], secondSubstring)

		substrings = append(substrings, substring)
	}

	tail := len(firstHalf) % maxLength
	wrote := len(firstHalf) - tail

	lastSubstring := make([]rune, tail*2)

	copy(lastSubstring, firstHalf[wrote:])
	copy(lastSubstring[tail:], secondHalf[:len(secondHalf)-wrote])

	substrings = append(substrings, lastSubstring)

	return substrings
}

type pool struct {
	workerNumber int
	workers      []worker
	took         time.Duration
	substrings   [][]rune
	result       bool
}

func (p *pool) poolCall() bool {
	sCh := make(chan []rune, len(p.substrings))
	resCh := make(chan bool, len(p.substrings))
	stopCh := make(chan interface{})

	if p.workerNumber > len(p.substrings) {
		p.workerNumber = len(p.substrings)
	}

	for i := 0; i < p.workerNumber; i++ {
		w := worker{
			id:         i,
			substrings: sCh,
			results:    resCh,
			stop:       stopCh,
		}

		p.workers = append(p.workers, w)

		go p.workers[i].work()
	}

	start := time.Now()

	for _, val := range p.substrings {
		sCh <- val
	}

	close(sCh)

	for i := 0; i < len(p.substrings); i++ {
		res := <-resCh
		if res == false {
			p.result = false

			break
		}
	}

	p.took = time.Since(start)

	close(stopCh)
	close(resCh)

	return p.result
}

type worker struct {
	id         int
	substrings chan []rune
	results    chan bool
	stop       chan interface{}
}

// work check is a substring is a palindrome.
func (w worker) work() {
	log.Printf("worker %d is up", w.id)

	for substring := range w.substrings {
		log.Printf("worker %d is checking if \"%s\" is a palindrome", w.id, string(substring))

		res := true

		for i := 0; i < len(substring)/2; i++ {
			r1 := unicode.ToLower(substring[i])
			r2 := unicode.ToLower(substring[len(substring)-i-1])

			if r1 != r2 {
				res = false
				break
			}
		}

		select {
		case <-w.stop:
			break
		default:
			w.results <- res
		}
	}
}

// write puts line in output.
func write(s string, out *bufio.Writer) error {
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

// read gets line from input and returns whole line and separated in two halves.
func read(in *bufio.Reader) (string, []rune, []rune, error) {
	b, _, err := in.ReadLine()
	if err != nil {
		return "", nil, nil, fmt.Errorf("read: %w", err)
	}

	word := string(b)

	if isAWord([]rune(word)) == false {
		return "", nil, nil, fmt.Errorf("examine: %w", errNotAWord)
	}

	firstHalf := []rune(string(b[0 : len(b)/2]))
	secondHalf := []rune(string(b[(len(b) - len(b)/2):]))

	return word, firstHalf, secondHalf, nil
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
