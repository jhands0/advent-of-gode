package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	total := 0
	var numbers [][]int
	var operators []string

	for _, line := range lines {
		if strings.Contains(line, "*") {
			break
		}
		for idx, str := range strings.Fields(line) {
			num, _ := strconv.Atoi(str)
			if idx >= len(numbers) {
				numbers = append(numbers, []int{})
			}
			numbers[idx] = append(numbers[idx], num)
		}
	}
	operators = append(operators, strings.Fields(lines[len(lines)-1])...)

	for i := range numbers {
		subtotal := 0
		switch operators[i] {
		case "*":
			subtotal = reduce(numbers[i], func(acc, current int) int {
				return acc * current
			})
		case "+":
			subtotal = reduce(numbers[i], func(acc, current int) int {
				return acc + current
			})
		default:
		}
		total += subtotal
	}

	fmt.Println(total)

}

func part2(lines []string) {
	total := 0
	var numbersStr [][]rune
	var operators []string

	for _, line := range lines {
		if strings.Contains(line, "*") {
			break
		}
		for idx, str := range []rune(line) {
			if idx >= len(numbersStr) {
				numbersStr = append(numbersStr, []rune{})
			}
			numbersStr[idx] = append(numbersStr[idx], str)
		}
	}
	operators = append(operators, strings.Fields(lines[len(lines)-1])...)

	var arr []int
	index := 0

	numbersStr = append(numbersStr, []rune(" "))
	for _, rune := range numbersStr {
		numStr := strings.TrimSpace(string(rune))
		subtotal := 0
		if numStr == "" {
			switch operators[index] {
			case "*":
				subtotal = reduce(arr, func(acc, current int) int {
					return acc * current
				})
			case "+":
				subtotal = reduce(arr, func(acc, current int) int {
					return acc + current
				})
			default:
			}
			arr = []int{}
			index++
			total += subtotal
			continue
		}
		num, _ := strconv.Atoi(numStr)
		arr = append(arr, num)
	}
	fmt.Println(total)
}

func reduce(arr []int, fn func(acc int, current int) int) int {
	acc := 0
	if len(arr) >= 1 {
		acc = arr[0]
	}
	for i := 1; i < len(arr); i++ {
		acc = fn(acc, arr[i])
	}
	return acc
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2025, 6)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
