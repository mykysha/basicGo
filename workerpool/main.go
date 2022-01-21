package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	var zones []int
	out := bufio.NewWriter(os.Stdout)

	for i := -11; i < 13; i++ {
		zones = append(zones, i)
	}

	p := pool{
		workerNumber: 4,
		zones:        zones,
	}

	p.poolCall()

	for _, val := range p.results {
		_, err := out.WriteString(val + "\n")
		if err != nil {
			log.Println(val, err)
		}

		err = out.Flush()
		if err != nil {
			log.Println(err)
		}
	}
}

type pool struct {
	workerNumber int
	workers      []worker
	zones        []int
	took         time.Duration
	results      []string
}

func (p *pool) poolCall() {
	zCh := make(chan int, len(p.zones))
	resCh := make(chan string, len(p.zones))

	for i := 0; i < p.workerNumber; i++ {
		w := worker{
			id:      i,
			zones:   zCh,
			results: resCh,
		}

		p.workers = append(p.workers, w)

		go p.workers[i].work()
	}

	start := time.Now()

	for i := 0; i < len(p.zones); i++ {
		zCh <- p.zones[i]
	}

	close(zCh)

	for i := 0; i < len(p.zones); i++ {
		res := <-resCh
		p.results = append(p.results, res)
	}

	p.took = time.Since(start)
}

type worker struct {
	id      int
	zones   chan int
	results chan string
}

func (w worker) work() {
	log.Printf("worker %d is up", w.id)

	for zone := range w.zones {
		log.Printf("worker %d is calculating time zone %d", w.id, zone)

		timeInZone := task(zone)

		zoneStr := strconv.Itoa(zone)

		if zone > 0 {
			zoneStr = "+" + zoneStr
		}

		res := fmt.Sprintf("Time in UTC%s is %v", zoneStr, timeInZone.Format("15:04 01.02"))

		w.results <- res
	}
}

func task(zone int) time.Time {
	if zone < -11 {
		zone = 1
	}

	if zone > 12 {
		zone = 24
	}

	timeInZone := time.Now().UTC().Add(time.Hour * time.Duration(zone))

	return timeInZone
}