package main

import (
	"advent-of-go/day3"
	"fmt"
	"os"
)

func main() {
	answer, error := day3.Run()

	if error != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %v\n", error)
	} else {
		fmt.Println(answer)
	}
}
