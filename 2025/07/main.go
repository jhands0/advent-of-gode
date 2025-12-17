package main

import (
	"fmt"
	"log"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	total := 0
	grid := [][]rune{}

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '^' && grid[i-1][j] == '|' {
				grid[i+1][j-1] = '|'
				grid[i+1][j+1] = '|'
				total += 1
			} else if grid[i][j] == '|' && i < len(grid)-1 && grid[i+1][j] == '.' {
				grid[i+1][j] = '|'
			} else if grid[i][j] == 'S' {
				grid[i+1][j] = '|'
				break
			}
		}
	}

	fmt.Print(grid)

	fmt.Println(total)
}

func part2(lines []string) {

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2025, 7)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
