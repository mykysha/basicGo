package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/nndergunov/basicGo/strings/anagram/anagramfinder"
	"github.com/nndergunov/basicGo/strings/anagram/cl"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	writer := cl.NewWriter(out)
	in := bufio.NewReader(os.Stdin)
	reader := cl.NewReader(in)
	a := anagramfinder.NewAnagramFinder()

	msg := "Please enter your first word:\t"

	err := writer(msg)
	if err != nil {
		log.Println(err)
	}

	firstWord, err := reader()
	if err != nil {
		log.Println(err)
	}

	msg = "Please enter your second word:\t"

	err = writer(msg)
	if err != nil {
		log.Println(err)
	}

	secondWord, err := reader()
	if err != nil {
		log.Println(err)
	}

	res := a.Compute(firstWord, secondWord)

	if res {
		msg = fmt.Sprintf("Words \"%s\" and \"%s\" are anagrams.", string(firstWord), string(secondWord))

		err = writer(msg)
		if err != nil {
			log.Println(err)
		}
	} else {
		msg = fmt.Sprintf("Words \"%s\" and \"%s\" are not anagrams.", string(firstWord), string(secondWord))

		err = writer(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
