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
	mult  = "*"
	add   = "+"
	space = " "
)

func Run() (int, error) {

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
	operators := strings.Fields(lines[operatorRow])
	operatorIndex := 0

	sanitizedLines := lo.Map(lines, func(s string, _ int) string { return strings.TrimRight(s, "\n") + space })

	maxLineWidth := lo.Max(lo.Map(sanitizedLines, func(s string, _ int) int { return len(s) }))

	problems := []problem{}
	numbers := []int{}

	for columnIndex := range maxLineWidth {

		currentNumberBuffer := strings.Builder{}

		for rowIndex := range operatorRow {

			if columnIndex >= len(sanitizedLines[rowIndex]) {
				continue
			}

			char := string(sanitizedLines[rowIndex][columnIndex])
			if char != space {
				currentNumberBuffer.WriteString(char)
			}
		}

		if "" == currentNumberBuffer.String() {
			problems = append(problems, problem{numbers, operators[operatorIndex]})
			numbers = []int{}
			operatorIndex++
			continue
		}

		intVal, error := strconv.Atoi(currentNumberBuffer.String())

		if error != nil {
			return nil, fmt.Errorf("Error parsing string: %v", error)
		}

		numbers = append(numbers, intVal)
	}

	return problems, nil
}
