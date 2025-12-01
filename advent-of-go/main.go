package main

import (
	"advent-of-go/day1"
	"fmt"
	"os"
)

func main() {
	answer, error := day1.Run()

	if error != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %s\n", error.Error())
	} else {
		fmt.Println(answer)
	}
}
