package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/echojc/aocutil"
)

func part1(lines []string) {
	var left []int
	var right []int

	for _, line := range lines {
		parts := strings.Fields(line)

		leftNum, _ := strconv.Atoi(parts[0])
		rightNum, _ := strconv.Atoi(parts[1])

		left = append(left, leftNum)
		right = append(right, rightNum)

	}

	sort.Ints(left)
	sort.Ints(right)

	totalDistance := 0

	for i := 0; i < len(left); i++ {
		distance := left[i] - right[i]

		if distance < 0 {
			distance *= -1
		}

		totalDistance += distance
	}

	fmt.Println(totalDistance)
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 1)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
}
