package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	records := make([][]int, len(lines))

	for i, line := range lines {
		strRecord := strings.Split(line, " ")
		intRecord := make([]int, len(strRecord))

		for j, val := range strRecord {
			numVal, _ := strconv.Atoi(val)
			intRecord[j] = numVal
		}

		records[i] = intRecord
	}

	safeRecords := 0

	for _, record := range records {
		prev := record[0]
		increasing := true
		safe := true

		if record[0] == record[1] {
			safe = false
		}
		if record[0] > record[1] {
			increasing = false
		}

		for i := 1; i < len(record); i++ {
			curr := record[i]

			if increasing && curr < prev {
				safe = false
			}
			if !increasing && curr > prev {
				safe = false
			}
			diff := prev - curr
			if diff < 0 {
				diff *= -1
			}
			if diff > 3 || diff < 1 {
				safe = false
			}
			prev = curr
		}

		if safe {
			safeRecords += 1
		}
	}

	fmt.Println(safeRecords)

}

func part2(lines []string) {

	fmt.Println(safeRecords)
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 2)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
