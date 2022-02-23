package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nndergunov/basicGo/slices/goqueue/printer"
	"github.com/nndergunov/basicGo/slices/goqueue/queue"
)

var errUnexpectedType = errors.New("did not expect to get type: ")

func main() {
	pl := log.New(os.Stdout, "printer ", log.LstdFlags)
	l := log.New(os.Stdout, "main ", log.LstdFlags)

	p := printer.NewPrinter(pl)

	var documents []string

	numberOfDocuments := 5

	for i := 0; i < numberOfDocuments; i++ {
		documents = append(documents, fmt.Sprintf("document%d.txt", i))
	}

	q := queue.NewQueue()

	for _, val := range documents {
		document := val

		q.Enqueue(document)
	}

	queueLen := q.GetLength()

	for i := 0; i < queueLen; i++ {
		task, err := q.Dequeue()
		if err != nil {
			l.Print(err)
		}

		doc, ok := task.(string)
		if !ok {
			l.Print(errUnexpectedType)
		}

		p.Print(doc)
	}
}
