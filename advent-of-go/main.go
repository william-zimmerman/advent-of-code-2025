package main

import (
	"advent-of-go/day8"
	"fmt"
	"os"
)

func main() {
	answer, err := day8.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Encountered err: %v\n", err)
	} else {
		fmt.Println(answer)
	}
}
