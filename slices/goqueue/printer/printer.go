package printer

import (
	"log"
	"time"
)

type Printer struct {
	l         *log.Logger
	printTime time.Duration
}

func NewPrinter(l *log.Logger) Printer {
	printSeconds := 3

	printTime := time.Second * time.Duration(printSeconds)

	return Printer{
		l:         l,
		printTime: printTime,
	}
}

func (p Printer) Print(document string) {
	time.Sleep(p.printTime)

	p.l.Printf("document %s had been printed", document)
}
