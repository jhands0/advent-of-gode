package main

import (
	"fmt"
	"log"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	count := 0
	var grid [][]rune

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	lenX := len(grid)
	lenY := len(grid[0])

	for row := range grid {
		for col, char := range grid[row] {

			if char == '@' {
				adj := 0

				for i := -1; i <= 1; i++ {
					newRow := row + i
					if newRow < 0 || newRow > lenX-1 {
						continue
					}

					for j := -1; j <= 1; j++ {
						newCol := col + j
						if i == 0 && j == 0 {
							continue
						}
						if newCol < 0 || newCol > lenY-1 {
							continue
						}
						if grid[newRow][newCol] == '@' {
							adj++
						}
					}
				}

				if adj < 4 {
					count++
				}
			}
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

	lines, err := input.Strings(2025, 4)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
