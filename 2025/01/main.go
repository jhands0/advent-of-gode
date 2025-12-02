package main

import (
	"fmt"
	"log"
	"strconv"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	var direction byte

	value := 50
	count := 0

	for _, line := range lines {

		direction = line[0]
		turn, _ := strconv.Atoi(line[1:])
		step := 1

		if direction == 'L' {
			step = 99
		}

		for range turn {
			value = (value + step) % 100
		}
		if value == 0 {
			count += 1
		}

	}

	fmt.Println(count)
}

func part2(lines []string) {
	var direction byte

	value := 50
	clicks := 0
	count := 0

	for _, line := range lines {

		direction = line[0]
		turn, _ := strconv.Atoi(line[1:])
		step := 1

		if direction == 'L' {
			step = 99
		}

		for range turn {
			value = (value + step) % 100
			if value == 0 {
				clicks += 1
			}
		}
		if value == 0 {
			count += 1
		}

	}

	fmt.Println(count, clicks)
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2025, 1)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
