package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nndergunov/basicGo/slices/goqueue/queue"
)

func main() {
	var documents []string

	numberOfDocuments := 5
	l := log.New(os.Stdout, "printer ", log.LstdFlags)

	for i := 0; i < numberOfDocuments; i++ {
		documents = append(documents, fmt.Sprintf("document%d.txt", i))
	}

	q := queue.NewQueue()

	for _, val := range documents {
		document := val

		toPrint := func() {
			printSeconds := 3

			printTime := time.Second * time.Duration(printSeconds)

			time.Sleep(printTime)

			l.Printf("document %s had been printed", document)
		}

		q.Enqueue(toPrint)
	}

	queueLen := q.GetLength()

	for i := 0; i < queueLen; i++ {
		task, err := q.Dequeue()
		if err != nil {
			l.Print(err)
		}

		task()
	}
}
