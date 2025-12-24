package main

import (
	"advent-of-go/day11"
	"fmt"
	"os"
)

func main() {
	answer, err := day11.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Encountered err: %v\n", err)
	} else {
		fmt.Println(answer)
	}
}
