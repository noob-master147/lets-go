package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	question string
	answer   string
}

func main() {
	fmt.Println("Welcome to the Quiz Game!")
	csvFile := flag.String("csv", "problem.csv", "Enter the path of the csv file in format of 'question,answer'.")
	timeLimit := flag.Int("limit", 30, "The time limit of the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file: %s", *csvFile))
	}
	r := csv.NewReader(file)
	// Extract the file
	lines, err := r.ReadAll()
	if err != nil {
		exit("Exit: Failed to parse the CSV file")
	}
	// parse the problems
	problems := parseLines(lines)
	count := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for i, p := range problems {
		answerChan := make(chan string)
		fmt.Printf("%d. %s = \n", i+1, p.question)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerChan:
			if answer == p.answer {
				count++
				fmt.Println("Correct !")
			}
		}

	}

	fmt.Printf("You scored %d/%d\n", count, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func parseLines(lines [][]string) []problem {
	r := make([]problem, len(lines))

	for i, line := range lines {
		r[i].question = line[0]
		r[i].answer = strings.TrimSpace(line[1])
	}
	return r
}
