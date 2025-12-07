package day7

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type matrix map[coordinate]string

type coordinate struct {
	x, y int
}

func (c coordinate) north() coordinate {
	return coordinate{c.x, c.y - 1}
}

func (c coordinate) east() coordinate {
	return coordinate{c.x + 1, c.y}
}

func (c coordinate) south() coordinate {
	return coordinate{c.x, c.y + 1}
}

func (c coordinate) west() coordinate {
	return coordinate{c.x - 1, c.y}
}

const (
	start    = "S"
	empty    = "."
	manifold = "^"
	beam     = "|"
)

func Run() (int, error) {

	fmt.Println("Hello from day7!")

	bytes, error := os.ReadFile("day7/input.txt")
	if error != nil {
		return 0, fmt.Errorf("Error reading file: %w", error)
	}

	matrix := parseInput(string(bytes))

	manifoldsEncountered, error := applyTaychonBeams(matrix)

	if error != nil {
		return 0, fmt.Errorf("Error encountered applying beams: %w", error)
	}

	return manifoldsEncountered, nil
}

func parseInput(fileContents string) matrix {

	matrix := matrix{}

	lines := slices.Collect(strings.Lines(fileContents))

	for row, line := range lines {
		sanitizedLine := strings.Fields(line)
		if len(sanitizedLine) != 1 {

		}

		for column, rune := range sanitizedLine[0] {
			coordinate := coordinate{column, row}
			matrix[coordinate] = string(rune)
		}
	}

	return matrix
}

func applyTaychonBeams(m matrix) (int, error) {

	entrance, error := manifoldEntrace(m)
	manifoldsEncountered := 0

	if error != nil {
		return 0, error
	}

	queue := []coordinate{}
	queue = append(queue, entrance.south())

	for len(queue) > 0 {
		poppedCoordinate := queue[0]
		queue = queue[1:]

		coordinateValue, found := m[poppedCoordinate]

		if !found {
			continue
		}

		switch coordinateValue {
		case empty:
			m[poppedCoordinate] = beam
			queue = append(queue, poppedCoordinate.south())

		case manifold:
			manifoldsEncountered++
			queue = append(queue, poppedCoordinate.east(), poppedCoordinate.west())

		case beam:
			continue

		default:
			return 0, fmt.Errorf("Encountered unexpected value %q at coordinate %v", coordinateValue, poppedCoordinate)
		}
	}

	return manifoldsEncountered, nil
}

func manifoldEntrace(matrix matrix) (coordinate, error) {
	startingCoordinate := coordinate{0, 0}

	var startCoordinateRecurse func(c coordinate) (coordinate, error)
	startCoordinateRecurse = func(c coordinate) (coordinate, error) {
		value, found := matrix[c]

		if !found {
			return coordinate{}, fmt.Errorf("Unable to find starting position")
		}

		switch value {
		case empty:
			return startCoordinateRecurse(c.east())
		case start:
			return c, nil
		}

		return coordinate{}, fmt.Errorf("Unexpected value encountered %q", value)
	}

	return startCoordinateRecurse(startingCoordinate)
}
