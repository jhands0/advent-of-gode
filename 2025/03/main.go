package main

import (
	"fmt"
	"log"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	total := 0

	for _, bank := range lines {
		batA := 0
		batB := 0

		for idx, r := range bank {
			num := int(r - '0')

			if num > batA && idx != len(bank)-1 {
				batA = num
				batB = 0
			} else if num > batB {
				batB = num
			}
		}
		total += batA*10 + batB
	}

	fmt.Println(total)
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
