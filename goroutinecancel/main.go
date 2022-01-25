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
	out := bufio.NewWriter(os.Stdout)
	in := bufio.NewReader(os.Stdin)

	l.Println("input/output and logger configured")

	var (
		wasAliveFor time.Duration
		reachedEnd  bool
	)

	signal := make(chan interface{})

	go routine(&wasAliveFor, &reachedEnd, signal)
	l.Println("routine started")

	err := waitForKill(in, out)
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

	err = write(msg, out)
	if err != nil {
		l.Println(err)
	}

	l.Println("end of program execution")
}

// routine counts how long it was alive.
func routine(isAliveFor *time.Duration, e *bool, s chan interface{}) {
	var sleepTime time.Duration = 100

	start := time.Now()

	for {
		select {
		case <-s:
			break
		default:
			*isAliveFor = time.Since(start)

			*e = true

			time.Sleep(sleepTime * time.Millisecond)
		}
	}
}

func waitForKill(in *bufio.Reader, out *bufio.Writer) error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("making terminal raw: %w", err)
	}

	defer func(fd int, oldState *term.State) error {
		err = term.Restore(fd, oldState)
		if err != nil {
			return fmt.Errorf("restoring usual terminal: %w", err)
		}

		return nil
	}(int(os.Stdin.Fd()), oldState)

	err = write("press any key to kill goroutine\n", out)
	if err != nil {
		return fmt.Errorf("prompting to kill: %w", err)
	}

	_, err = in.ReadByte()
	if err != nil {
		return fmt.Errorf("reading info: %w", err)
	}

	return nil
}

func write(msg string, out *bufio.Writer) error {
	_, err := out.WriteString(msg)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	err = out.Flush()
	if err != nil {
		return fmt.Errorf("flush: %w", err)
	}

	return nil
}
