package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	var separator int

	for idx, line := range lines {
		if line == "" {
			separator = idx
		}
	}

	rules := lines[:separator]
	ruleMap := parseRules(rules)
	updates := lines[separator+1:]
	updateList := parseUpdates(updates)

	result := 0
	for _, update := range updateList {
		if checkIfValid(update, ruleMap) {
			result += update[len(update)/2]
		}
	}

	fmt.Println(result)
}

func parseRules(rules []string) map[int][]int {
	ruleMap := make(map[int][]int)

	for _, rule := range rules {
		splitRule := strings.Split(rule, "|")
		key, _ := strconv.Atoi(splitRule[0])
		value, _ := strconv.Atoi(splitRule[1])

		ruleMap[key] = append(ruleMap[key], value)
	}

	return ruleMap
}

func parseUpdates(updates []string) [][]int {
	var updateList [][]int

	for _, update := range updates {
		updateSplit := strings.Split(update, ",")
		var updateSplitInt []int

		for _, val := range updateSplit {
			convertedVal, _ := strconv.Atoi(val)
			updateSplitInt = append(updateSplitInt, convertedVal)
		}

		updateList = append(updateList, updateSplitInt)
	}

	return updateList
}

func checkIfValid(update []int, rulesMap map[int][]int) bool {
	for idx := 1; idx < len(update); idx++ {
		for _, val := range rulesMap[update[idx]] {
			if val == update[idx-1] {
				return false
			}
		}
	}

	return true
}

func part2(lines []string) {
	var separator int

	for idx, line := range lines {
		if line == "" {
			separator = idx
		}
	}

	rules := lines[:separator]
	ruleMap := parseRules(rules)
	updates := lines[separator+1:]
	updateList := parseUpdates(updates)

	result := 0
	for _, update := range updateList {
		if !checkIfValid(update, ruleMap) {
			fixedUpdate := fixUpdate(update, ruleMap)
			result += fixedUpdate[len(update)/2]
		}
	}

	fmt.Println(result)
}

func fixUpdate(update []int, rulesMap map[int][]int) []int {
	fixedUpdate := make([]int, len(update))

	for i := 0; i < len(update); i++ {
		countFound := 0
		if _, ok := rulesMap[update[i]]; !ok {
			countFound = 0
		} else {
			for j := 0; j < len(update); j++ {
				if i == j {
					continue
				}
				found := false
				for _, val := range rulesMap[update[i]] {
					if val == update[j] {
						found = true
						break
					}
				}

				if found {
					countFound += 1
				}
			}
		}

		fixedUpdate[len(update)-countFound-1] = update[i]
	}

	return fixedUpdate
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 5)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
