package main

import (
	"advent-of-go/day2"
	"fmt"
	"os"
)

func main() {
	answer, error := day2.Run()

	if error != nil {
		fmt.Fprintf(os.Stderr, "Encountered error: %s\n", error.Error())
	} else {
		fmt.Println(answer)
	}
}
