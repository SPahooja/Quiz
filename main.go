package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile := flag.String("csv", "sample.csv", "csv file with question,answer")
	flag.Parse()

	file, err := os.Open(*csvFile)

	if err != nil {
		exit(fmt.Sprintf("Unable to open %s\n", *csvFile))
	}
	read := csv.NewReader(file)
	lines, err := read.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse %s\n", *csvFile))
	}
	problems := problemParser(lines)
	fmt.Println(problems)
}

type Question struct {
	q string
	a string
}

func problemParser(lines [][]string) []Question {
	problem := make([]Question, len(lines))
	for i, line := range lines {
		problem[i] = Question{
			q: line[0],
			a: line[1],
		}
	}
	return problem
}

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
