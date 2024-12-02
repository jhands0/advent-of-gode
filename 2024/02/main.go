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
		if isSafe(record) {
			safeRecords += 1
		}
	}

	fmt.Println(safeRecords)

}

func part2(lines []string) {
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
		if !isSafe(record) {
			if bruteForceRecord(record) {
				safeRecords += 1
			}
		} else {
			safeRecords += 1
		}
	}

	fmt.Println(safeRecords)
}

func isSafe(record []int) bool {
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

	return safe
}

func bruteForceRecord(record []int) bool {
	for i := 0; i < len(record); i++ {
		copyRecord := make([]int, len(record))
		copy(copyRecord, record)

		if i == len(copyRecord)-1 {
			copyRecord = copyRecord[:i]
		} else {
			copyRecord = append(copyRecord[:i], copyRecord[i+1:]...)
		}

		if isSafe(copyRecord) {
			return true
		}
	}

	return false
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
