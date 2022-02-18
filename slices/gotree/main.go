package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nndergunov/basicGo/slices/gotree/binarytree"
)

func main() {
	l := log.New(os.Stdout, "penguin tree ", log.LstdFlags)
	writer := newWriter()

	species := []string{
		"King", "Emperor", "Adélie", "Chinstrap", "Gentoo", "Little", "Australian little",
		"White-flippered", "Magellanic", "Humboldt", "Galapagos", "African", "Yellow-eyed",
		"Fiordland", "Snares", "Erect-crested", "Southern rockhopper", "Eastern rockhopper",
		"Northern rockhopper", "Royal", "Macaroni",
	}

	penguins := binarytree.NewBinaryNode(species[0])

	for _, val := range species {
		err := penguins.Insert(val)

		if err != nil && !errors.Is(err, binarytree.ErrNodeExists) {
			l.Println(err)
		}
	}

	lowestValuePenguin := penguins.GetMinValue()

	err := writer(fmt.Sprintf("\nThe most-left node is %s penguin\n", lowestValuePenguin))
	if err != nil {
		log.Println(err)
	}

	highestValuePenguin := penguins.GetMaxValue()

	err = writer(fmt.Sprintf("\nThe most-right node is %s penguin\n", highestValuePenguin))
	if err != nil {
		log.Println(err)
	}

	orderedPenguins := penguins.GetInOrder()

	err = writer("\nPenguins in order of the binary tree:\n")
	if err != nil {
		log.Println(err)
	}

	for _, val := range orderedPenguins {
		err = writer(fmt.Sprintf("%s penguin\n", val))
		if err != nil {
			log.Println(err)
		}
	}
}

func newWriter() func(s string) error {
	w := bufio.NewWriter(os.Stdout)

	return func(s string) error {
		_, err := w.WriteString(s)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}

		err = w.Flush()
		if err != nil {
			return fmt.Errorf("flush: %w", err)
		}

		return nil
	}
}
