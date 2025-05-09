package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvFile := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFile))
	}
	defer file.Close()

	fileReader := csv.NewReader(file)
	records, err := fileReader.ReadAll()

	problems := getProblems(records)
	// fmt.Println(problems)

	correct := 0
	for i, question := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, question.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == question.a {
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d..", correct, len(problems))

}

func getProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, prob := range lines {
		problems[i] = problem{
			q: prob[0],
			a: strings.TrimSpace(prob[1]),
		}
	}

	return problems
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
