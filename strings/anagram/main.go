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

	msg := "\nPlease enter your first word:\t"

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

	res, err := anagramfinder.CheckIfAnagram(firstWord, secondWord)
	if err != nil {
		log.Println(err)
	}

	msg = fmt.Sprintf("Words \"%s\" and \"%s\" are", firstWord, secondWord)

	if !res {
		msg += " not"
	}

	msg += " anagrams\n"

	err = writer(msg)
	if err != nil {
		log.Println(err)
	}
}
