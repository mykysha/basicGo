package printer

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nndergunov/basicGo/slices/goqueue/queue"
)

var errUnexpectedType = errors.New("did not expect to get this type")

type Printer struct {
	l         *log.Logger
	printTime time.Duration
	queue     *queue.SliceQueue
}

func NewPrinter(l *log.Logger) Printer {
	printSeconds := 3

	printTime := time.Second * time.Duration(printSeconds)

	q := queue.NewQueue()

	return Printer{
		l:         l,
		printTime: printTime,
		queue:     &q,
	}
}

func (p *Printer) Add(docs []string) {
	for _, val := range docs {
		document := val

		p.queue.Enqueue(document)
	}
}

func (p *Printer) PrintQueue() error {
	docNum := p.queue.GetLength()

	for i := 0; i < docNum; i++ {
		task, err := p.queue.Dequeue()
		if err != nil {
			p.l.Print(err)

			return fmt.Errorf("dequeue document: %w", err)
		}

		doc, ok := task.(string)
		if !ok {
			p.l.Print(errUnexpectedType)

			return errUnexpectedType
		}

		p.PrintDoc(doc)
	}

	return nil
}

func (p Printer) PrintDoc(document string) {
	time.Sleep(p.printTime)

	p.l.Printf("document %s had been printed", document)
}
