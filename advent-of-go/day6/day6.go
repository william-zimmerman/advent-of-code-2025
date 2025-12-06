package day6

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/lo/it"
)

type problem struct {
	input     []int
	operation string
}

func (p problem) solve() int {
	switch p.operation {
	case add:
		return lo.Sum(p.input)
	case mult:
		return lo.Product(p.input)
	}
	panic(fmt.Sprintf("Unexpected operation %q", p.operation))
}

const (
	mult = "*"
	add  = "+"
)

func Run() (int, error) {

	fmt.Println("Hello from day6!")

	bytes, error := os.ReadFile("day6/input.txt")
	if error != nil {
		return 0, fmt.Errorf("Error parsing file: %v", error)
	}

	lines := slices.Collect(strings.Lines(string(bytes)))
	problems, error := parseInput(lines)

	if error != nil {
		return 0, fmt.Errorf("Error parsing input: %v", error)
	}

	sum := it.Sum(it.Map(slices.Values(problems), func(p problem) int { return p.solve() }))

	return sum, nil
}

func parseInput(lines []string) ([]problem, error) {

	operatorRow := len(lines) - 1
	matrix := slices.Collect(it.Map(slices.Values(lines), strings.Fields))
	columnCount := len(matrix[0])

	problems := []problem{}
	for columnIndex := range columnCount {
		problemInput := []int{}
		for rowIndex := range operatorRow {

			intVal, error := strconv.Atoi(matrix[rowIndex][columnIndex])

			if error != nil {
				return nil, error
			}

			problemInput = append(problemInput, intVal)
		}

		problems = append(problems, problem{problemInput, matrix[operatorRow][columnIndex]})
	}

	return problems, nil
}
