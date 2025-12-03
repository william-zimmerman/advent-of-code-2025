package main

import (
	"advent-of-go/day3"
	"fmt"
	"os"
)

func main() {
	answer, error := day3.Run()

	if error != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %s\n", error.Error())
	} else {
		fmt.Println(answer)
	}
}
