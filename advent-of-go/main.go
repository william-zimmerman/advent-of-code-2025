package main

import (
	"advent-of-go/day5"
	"fmt"
	"os"
)

func main() {
	answer, error := day5.Run()

	if error != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %v\n", error)
	} else {
		fmt.Println(answer)
	}
}
