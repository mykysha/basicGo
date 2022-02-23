package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nndergunov/basicGo/slices/goqueue/printer"
)

func main() {
	pl := log.New(os.Stdout, "printer ", log.LstdFlags)
	l := log.New(os.Stdout, "main ", log.LstdFlags)
	p := printer.NewPrinter(pl)

	var documents []string

	numberOfDocuments := 5

	for i := 0; i < numberOfDocuments; i++ {
		documents = append(documents, fmt.Sprintf("document%d.txt", i))
	}

	p.Add(documents)

	if err := p.PrintQueue(); err != nil {
		l.Println(err)
	}
}
