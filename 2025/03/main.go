package main

import (
	"log"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {

}

func part2(lines []string) {

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2025, 3)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
