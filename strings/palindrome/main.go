package main

import (
	"bufio"
	"fmt"
	"github.com/nndergunov/basicGo/strings/palindrome/cl"
	"github.com/nndergunov/basicGo/strings/palindrome/palindromefinder"
	"log"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	writer := cl.NewWriter(out)
	in := bufio.NewReader(os.Stdin)
	reader := cl.NewReader(in)

	msg := fmt.Sprint("\nThis program will tell if your word is a palindrome\n" +
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

	res := fmt.Sprintf("\nThe word \"%s\" is", word)

	result, err := palindromefinder.RunConcurrently(word)
	if err != nil {
		log.Println(err)
	}

	if !result || err != nil {
		res += " not"
	}

	res += " a palindrome\n"

	err = writer(res)
	if err != nil {
		log.Println(err)
	}
}
