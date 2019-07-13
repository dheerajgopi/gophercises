package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Problem statement and solution
type Problem struct {
	statement string
	solution  string
}

// parseCliArgs parse parameters passed to the application through the console.
// Default values are applied if arguments are not provided.
func parseCliArgs() (string, bool, int) {
	var csvFilename string
	var random bool
	var timeout int

	flag.StringVar(
		&csvFilename,
		"questions",
		"problems.csv",
		"Specify the csv file containing the questions and answers")
	flag.BoolVar(
		&random,
		"random",
		false,
		"randomize order of the questions",
	)
	flag.IntVar(
		&timeout,
		"timeout",
		30,
		"timeout in seconds",
	)

	flag.Parse()

	return csvFilename, random, timeout
}

// readAnswer waits for user input from the console
func readAnswer(answer chan string) {
	var input string
	fmt.Scanln(&input)

	answer <- input
}

// checkTimeout sleeps for a set time (in seconds) and then
// sets a boolean status into a channel
func checkTimeout(seconds int, timeout chan bool) {
	time.Sleep(time.Duration(seconds) * time.Second)

	timeout <- true
}

// printSummary prints the final score
func printSummary(score int, total int) {
	fmt.Println(fmt.Sprintf("Scored %d out of %d!!", score, total))
}

func main() {
	fileName, randomOrder, timeoutSeconds := parseCliArgs()
	csvFile, err := os.Open(fileName)

	if err != nil {
		log.Fatal(fmt.Sprintf("failed to open the csv file: %s\n", fileName))
	}

	defer csvFile.Close()

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
	total := len(problems)

	// randomize question order if 'random' flag is set
	questionOrder := make([]int, total)

	if randomOrder {
		rand.Seed(time.Now().UnixNano())
		questionOrder = rand.Perm(total)
	} else {
		for i := range questionOrder {
			questionOrder[i] = i
		}
	}

	timeout := make(chan bool, 1)
	answerInput := make(chan string, 1)

	// start quiz on key press
	fmt.Println("Press [Enter] to start quiz")
	bufio.NewScanner(os.Stdout).Scan()

	// check for timeout
	go checkTimeout(timeoutSeconds, timeout)

	// print a question and wait for user input from the console.
	// if user enters the right answer, increment the score.
	// quiz will stop after timeout
	for _, qtnNum := range questionOrder {
		fmt.Print(problems[qtnNum].statement + ": ")

		go readAnswer(answerInput)

		select {
		case answer := <-answerInput:
			if strings.ToLower(strings.TrimSpace(answer)) == problems[qtnNum].solution {
				score++
			}
		case <-timeout:
			fmt.Println("TIMEOUT!!")
			printSummary(score, total)
			return
		}

	}

	printSummary(score, total)
}
