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
	freshIngredients := make(map[int]int)
	var ingredients []string

	for i, line := range lines {
		if line == "" {
			ingredients = lines[i:]
			break
		}

		split := strings.Split(line, "-")
		low, _ := strconv.Atoi(split[0])
		up, _ := strconv.Atoi(split[1])

		if value, found := freshIngredients[low]; found {
			if value > up {
				continue
			}
		}
		freshIngredients[low] = up
	}

	for _, ingredientStr := range ingredients {
		ingredient, _ := strconv.Atoi(ingredientStr)
		for low, high := range freshIngredients {
			if ingredient >= low && ingredient <= high {
				total += 1
				break
			}
		}
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

	lines, err := input.Strings(2025, 5)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
