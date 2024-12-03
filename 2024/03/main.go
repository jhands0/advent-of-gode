package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	exp, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	nums, _ := regexp.Compile(`[0-9]+`)

	var parsedLine []string

	for _, line := range lines {
		parsedLine = append(parsedLine, exp.FindAllString(line, -1)...)
	}

	total := 0

	for _, operation := range parsedLine {
		values := nums.FindAllString(operation, -1)

		leftNum, _ := strconv.Atoi(values[0])
		rightNum, _ := strconv.Atoi(values[1])

		total += leftNum * rightNum
	}

	fmt.Println(total)

}

func part2(lines []string) {
	exp, _ := regexp.Compile(`mul\([0-9]+,[0-9]+\)|do\(\)|don't\(\)`)
	nums, _ := regexp.Compile(`[0-9]+`)

	var parsedLines []string

	for _, line := range lines {
		parsedLines = append(parsedLines, exp.FindAllString(line, -1)...)
	}

	total := 0
	disabled := false

	for _, parsedLine := range parsedLines {
		if parsedLine == "do()" {
			disabled = false
		} else if parsedLine == "don't()" {
			disabled = true
		} else {
			if disabled {
				continue
			} else {
				values := nums.FindAllString(parsedLine, -1)

				leftNum, _ := strconv.Atoi(values[0])
				rightNum, _ := strconv.Atoi(values[1])

				total += leftNum * rightNum
			}
		}
	}

	fmt.Println(total)
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 3)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
