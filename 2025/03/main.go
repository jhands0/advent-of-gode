package main

import (
	"fmt"
	"log"
	"math"

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
	total := 0

	for _, bank := range lines {
		batteries := [12]int{}
		lenBat := len(batteries)
		lenBank := len(bank)

		for idx, r := range bank {
			num := int(r - '0')
			var start int
			if start = lenBat - (lenBank - idx); start < 0 {
				start = 0
			}
			for i := start; i < lenBat; i++ {
				if num > batteries[i] {
					batteries[i] = num
					for j := i + 1; j < lenBat; j++ {
						batteries[j] = 0
					}
					break
				}
			}
		}

		for i, c := range batteries {
			mul := lenBat - i - 1
			total += c * int(math.Pow10(mul))
		}
	}

	println(total)
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
