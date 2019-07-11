package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Problem statement and solution
type Problem struct {
	statement string
	solution  string
}

// parseCliArgs parse parameters passed to the application through the console.
// Default values are applied if arguments are not provided.
func parseCliArgs() string {
	var csvFilename string

	flag.StringVar(
		&csvFilename,
		"questions",
		"problems.csv",
		"Specify the csv file containing the questions and answers")

	flag.Parse()

	return csvFilename
}

func main() {
	fileName := parseCliArgs()
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(fmt.Sprintf("failed to open the csv file: %s\n", fileName))
	}

	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	var problems []Problem

	// Parse questions and answers from csv file
	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("error reading the csv file")
		}

		problems = append(problems, Problem{
			statement: strings.ToLower(strings.TrimSpace(line[0])),
			solution:  strings.ToLower(strings.TrimSpace(line[1])),
		})
	}

	score := 0
	totalScore := len(problems)

	// print a question and wait for user input from the console.
	// if user enters the right answer, increment the score.
	for _, p := range problems {
		fmt.Print(p.statement + ": ")
		var answer string
		fmt.Scanln(&answer)

		if strings.ToLower(strings.TrimSpace(answer)) == p.solution {
			score++
		}
	}

	fmt.Println(fmt.Sprintf("Scored %d out of %d!!", score, totalScore))
}
