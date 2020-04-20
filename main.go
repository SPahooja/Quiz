package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFile := flag.String("csv", "sample.csv", "csv file with question,answer")
	timelimit := flag.Int("limit", 10, "time limit in seconds")
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

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	correctCount := 0

	for i, p := range problems {
		fmt.Printf("Question #%d : %s\n", i+1, p.q)
		ansch := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansch <- answer
		}()
		/*reader := bufio.NewReader(os.Stdin)
		b := []byte("\n")
		answer, err := reader.ReadString(b[0])
		if err != nil {
			exit(fmt.Sprintf("Unable to read Answer"))
		}
		answer = strings.TrimSuffix(answer, "\n")
		fmt.Println(answer, p.a)*/
		select {
		case <-timer.C:
			fmt.Printf("\nTime limit of %d seconds ran out\n", *timelimit)
			fmt.Printf("You scored %d out of %d\n", correctCount, len(problems))
			return
		case answer := <-ansch:
			if answer == p.a {
				fmt.Println("Correct!")
				correctCount++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}
	fmt.Printf("You scored %d out of %d\n", correctCount, len(problems))
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
			a: strings.TrimSpace(line[1]),
		}
	}
	return problem
}

func exit(s string) {
	fmt.Println(s)
	os.Exit(1)
}
