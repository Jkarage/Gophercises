package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problems struct {
	q string
	a string
}

func main() {
	filename := flag.String("filename", "problems.csv", "A problems file in csv format")
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
	var correct int
	for i, v := range p {
		var answer string
		fmt.Printf("Problem No %d: %v\n", i, v.q)
		fmt.Scanf("%s", &answer)
		if answer == v.a {
			correct++
		}
	}
	fmt.Printf("Got %d correct out of %v", correct, len(p))
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
