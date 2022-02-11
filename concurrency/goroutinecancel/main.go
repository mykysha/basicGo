package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/term"
)

// main func starts new goroutine and waits for user to stop its execution.
func main() {
	l := log.New(os.Stdout, "gc ", log.Ltime)
	l.Println("logger created")

	out := bufio.NewWriter(os.Stdout)
	l.Println("output configured")

	in := bufio.NewReader(os.Stdin)
	l.Println("input configured")

	var (
		wasAliveFor time.Duration
		reachedEnd  bool
	)

	signal := make(chan interface{})

	go routine(&wasAliveFor, &reachedEnd, signal)
	l.Println("routine started")

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		l.Fatalln(err)
	}

	defer func(fd int, oldState *term.State) {
		err = term.Restore(fd, oldState)
		if err != nil {
			l.Println(err)
		}
	}(int(os.Stdin.Fd()), oldState)

	_, err = out.WriteString("press any key to kill goroutine\n")
	if err != nil {
		l.Println(err)
	}

	err = out.Flush()
	if err != nil {
		l.Println(err)
	}

	_, err = in.ReadByte()
	if err != nil {
		l.Println(err)
	}

	l.Println("closing goroutine")

	close(signal)

	msg := fmt.Sprintf("goroutine was alive for %v\n", wasAliveFor)

	if reachedEnd {
		msg += "routine had reached its end\n"
	} else {
		msg += "routine had not reached its end\n"
	}

	_, err = out.WriteString(msg)
	if err != nil {
		l.Println(err)
	}

	err = out.Flush()
	if err != nil {
		l.Println(err)
	}

	l.Println("end of program execution")
}

// routine counts how long it was alive.
func routine(isAliveFor *time.Duration, e *bool, s chan interface{}) {
	start := time.Now()

	for {
		select {
		case <-s:
			break
		default:
			*isAliveFor = time.Since(start)

			*e = true

			time.Sleep(100 * time.Millisecond)
		}
	}
}
