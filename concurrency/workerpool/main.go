package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nndergunov/basicGo/concurrency/workerpool/pool"
)

func main() {
	var zones []int

	out := bufio.NewWriter(os.Stdout)
	writeLine := newWriter(out)

	firstTimeZone := -11
	lastTimeZone := 13

	for i := firstTimeZone; i < lastTimeZone; i++ {
		zones = append(zones, i)
	}

	workerNumber := 4

	p := pool.NewPool(workerNumber)

	p.Start()

	type result struct {
		zone       int
		timeInZone time.Time
	}

	resultC := make(chan result, len(zones))

	for _, val := range zones {
		z := val

		task := func() {
			timeInWorld := calculateWorldTime(z)

			resultC <- result{
				zone:       z,
				timeInZone: timeInWorld,
			}
		}

		p.AddTask(task)
	}

	for i := 0; i < len(zones); i++ {
		res := <-resultC
		zoneStr := strconv.Itoa(res.zone)
		timeStr := res.timeInZone.Format("15:04 01.02")
		resStr := fmt.Sprintf("Time in UTC%s is %v", zoneStr, timeStr)

		err := writeLine(resStr)
		if err != nil {
			log.Println(err)
		}
	}
}

func calculateWorldTime(zone int) time.Time {
	minimalZone := -11
	maximalZone := 12

	if zone < minimalZone {
		zone = minimalZone
	} else if zone > maximalZone {
		zone = maximalZone
	}

	timeInZone := time.Now().UTC().Add(time.Hour * time.Duration(zone))

	return timeInZone
}

func newWriter(out *bufio.Writer) func(s string) error {
	return func(s string) error {
		_, err := out.WriteString(s + "\n")
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
