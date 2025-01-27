package main

import (
	"log"
	"strconv"
	"strings"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	result := 0

	for _, line := range lines {
		line := strings.Split(line, " ")
		total := line[0][:len(line[0])-1]
		totalNum, _ := strconv.Atoi(total)

		var lineNum []int
		for idx := 1; idx < len(line); idx++ {
			lineNum[idx-1], _ = strconv.Atoi(line[idx])
		}

		//BACKTRACKING
		working := lineNum[0]

	}
}

func part2(lines []string) {

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 7)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
