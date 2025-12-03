package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	lines = strings.Split(lines[0], ",")
	count := 0

	for _, line := range lines {
		s := strings.Split(line, "-")
		lowerStr := s[0]
		upperStr := s[1]

		if len(lowerStr)%2 != 0 {
			lowerStr = "1" + strings.Repeat("0", len(lowerStr))
		}
		if len(upperStr)%2 != 0 {
			upperStr = "9" + strings.Repeat("9", len(upperStr)-1)
		}

		lower, _ := strconv.Atoi(lowerStr)
		upper, _ := strconv.Atoi(upperStr)
		if lower > upper {
			continue
		}

		lowerHalfStr := lowerStr[:len(lowerStr)/2]
		upperHalfStr := upperStr[:len(upperStr)/2]
		lowerHalf, _ := strconv.Atoi(lowerHalfStr)
		upperHalf, _ := strconv.Atoi(upperHalfStr)

		candidates := []int{}

		for i := lowerHalf; i <= upperHalf; i++ {
			candidateStr := strings.Repeat(strconv.Itoa(i), 2)
			candidate, _ := strconv.Atoi(candidateStr)

			if candidate >= lower && candidate <= upper {
				candidates = append(candidates, candidate)
			}
		}

		for _, c := range candidates {
			count += c
		}
	}

	fmt.Println(count)
}

func part2(lines []string) {

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2025, 2)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
