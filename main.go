package main

import (
	"fmt"
	"log"

	"github.com/echojc/aocutil"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}

	lines, err := input.Strings(2024, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lines)
}
