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
	writer := newWriter(bufio.NewWriter(os.Stdout))

	l.Println("output and logger configured")

	var (
		wasAliveFor time.Duration
		reachedEnd  bool
	)

	signal := make(chan interface{})

	go routine(&wasAliveFor, &reachedEnd, signal)
	l.Println("routine started")

	err := writer("press any key to kill goroutine\n")
	if err != nil {
		l.Println(err)
	}

	err = getKey()
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

	err = writer(msg)
	if err != nil {
		l.Println(err)
	}

	l.Println("end of program execution")
}

// routine counts how long it was alive.
func routine(isAliveFor *time.Duration, e *bool, s chan interface{}) {
	start := time.Now()

	sleepFor := 100

	sleepTime := time.Duration(sleepFor) * time.Millisecond

	for {
		select {
		case <-s:
			break
		default:
			*isAliveFor = time.Since(start)

			*e = true

			time.Sleep(sleepTime)
		}
	}
}

func getKey() error {
	in := bufio.NewReader(os.Stdin)

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("making raw: %w", err)
	}

	defer func(fd int, oldState *term.State) {
		err = term.Restore(fd, oldState)
		if err != nil {
			panic(fmt.Errorf("could not restore %w", err))
		}
	}(int(os.Stdin.Fd()), oldState)

	_, err = in.ReadByte()
	if err != nil {
		return fmt.Errorf("reading: %w", err)
	}

	return nil
}

func newWriter(out *bufio.Writer) func(s string) error {
	return func(s string) error {
		_, err := out.WriteString(s)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}

		err = out.Flush()
		if err != nil {
			return fmt.Errorf("flush: %w", err)
		}

		return nil
	}
}
