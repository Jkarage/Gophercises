package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problems struct {
	q string
	a string
}

func main() {
	filename := flag.String("f", "problems.csv", "A problems file in csv format")
	timer := flag.Int("t", 30, "A timer for the quiz")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	p := parseRecords(records)
	c := time.NewTimer(time.Duration(*timer) * time.Second)
	correct := start(p, c)
	fmt.Printf("Got %d out of %v", correct, len(p))
}

func parseRecords(a [][]string) []problems {
	r := make([]problems, len(a))
	for i, v := range a {
		r[i] = problems{
			q: v[0],
			a: v[1],
		}
	}
	return r
}

func start(p []problems, c *time.Timer) int {
	var correct int
	for i, v := range p {
		fmt.Printf("Problem No %d: %v\n", i, v.q)
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChan <- answer
		}()
		select {
		case <-c.C:
			return correct
		case answer := <-answerChan:
			if answer == v.a {
				correct++
			}
		}
	}
	return correct
}
