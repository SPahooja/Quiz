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
	fmt.Println(lines)
}

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
