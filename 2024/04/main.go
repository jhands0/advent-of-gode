package main

import (
	"fmt"
	"log"

	aocutil "github.com/echojc/aocutil"
)

func part1(lines []string) {
	count := 0

	for row := 0; row < len(lines); row++ {
		for col, char := range lines[row] {
			if char == 'X' {
				count += checkForWord(lines, row, col)
			}
		}
	}

	fmt.Println(count)
}

func checkForWord(grid []string, row int, col int) int {
	count := 0

	// Up-Left
	if row >= 3 && col >= 3 {
		if grid[row-1][col-1] == 'M' && grid[row-2][col-2] == 'A' && grid[row-3][col-3] == 'S' {
			count += 1
		}
	}

	// Up
	if row >= 3 {
		if grid[row-1][col] == 'M' && grid[row-2][col] == 'A' && grid[row-3][col] == 'S' {
			count += 1
		}
	}

	// Up-Right
	if row >= 3 && col <= (len(grid[0])-1-3) {
		if grid[row-1][col+1] == 'M' && grid[row-2][col+2] == 'A' && grid[row-3][col+3] == 'S' {
			count += 1
		}
	}

	// Right
	if col <= (len(grid[0]) - 1 - 3) {
		if grid[row][col+1] == 'M' && grid[row][col+2] == 'A' && grid[row][col+3] == 'S' {
			count += 1
		}
	}

	// Down-Right
	if row <= (len(grid[0])-1-3) && col <= (len(grid[0])-1-3) {
		if grid[row+1][col+1] == 'M' && grid[row+2][col+2] == 'A' && grid[row+3][col+3] == 'S' {
			count += 1
		}
	}

	// Down
	if row <= (len(grid[0]) - 1 - 3) {
		if grid[row+1][col] == 'M' && grid[row+2][col] == 'A' && grid[row+3][col] == 'S' {
			count += 1
		}
	}

	// Down-Left
	if row <= (len(grid[0])-1-3) && col >= 3 {
		if grid[row+1][col-1] == 'M' && grid[row+2][col-2] == 'A' && grid[row+3][col-3] == 'S' {
			count += 1
		}
	}

	// Left
	if col >= 3 {
		if grid[row][col-1] == 'M' && grid[row][col-2] == 'A' && grid[row][col-3] == 'S' {
			count += 1
		}
	}

	return count
}

func part2(lines []string) {
	count := 0

	for row := 0; row < len(lines); row++ {
		for col, char := range lines[row] {
			if char == 'A' {
				count += checkForCross(lines, row, col)
			}
		}
	}

	fmt.Println(count)
}

func checkForCross(grid []string, row int, col int) int {
	count := 0

	if row < 1 || row > len(grid[0])-2 || col < 1 || col > len(grid[0])-2 {
		return 0
	}

	// M.M
	// .A.
	// S.S
	if grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'M' && grid[row+1][col+1] == 'S' && grid[row+1][col-1] == 'S' {
		count += 1
	}

	// S.M
	// .A.
	// S.M
	if grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'M' && grid[row+1][col+1] == 'M' && grid[row+1][col-1] == 'S' {
		count += 1
	}

	// S.S
	// .A.
	// M.M
	if grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'S' && grid[row+1][col+1] == 'M' && grid[row+1][col-1] == 'M' {
		count += 1
	}

	// M.S
	// .A.
	// M.S
	if grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'S' && grid[row+1][col+1] == 'S' && grid[row+1][col-1] == 'M' {
		count += 1
	}

	return count
}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 4)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
