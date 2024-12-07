package main

import (
	"log"

	aocutil "github.com/echojc/aocutil"
)

type position struct {
	x int
	y int
}

func part1(lines []string) {
	var startY int
	var startX int

	for y, line := range lines {
		for x, val := range line {
			if val == '^' {
				startX = x
				startY = y
				break
			}
		}
	}

	seen := make(map[position]bool)

	guardPos := position{startX, startY}
    guardDir := position{0, -1}
    count := 0

	for {
        if guardPos.x == 0 || guardPos.x == len(lines[0]) - 1 {
            break
        }
        if guardPos.y == 0 || guardPos.y == len(lines) - 1 {
            break
        }

        nextX := guardPos.x + guardDir.x
		nextY := guardPos.y + guardDir.y
        nextPos := position{nextX, nextY}

        if _, ok := seen[nextPos]; ok {
            guardPos = nextPos
        }

        if lines[nextY][nextY] == '.' && ne
	}
}

func part2(lines []string) {

}

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 6)
	if err != nil {
		log.Fatal(err)
	}

	part1(lines)
	part2(lines)
}
