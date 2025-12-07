package day7

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type matrix map[coordinate]string
type pathCountCache map[coordinate]int

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

	pathCountCache := pathCountCache{}
	uniquePathCount, error := uniquePaths(matrix, pathCountCache)

	if error != nil {
		return 0, fmt.Errorf("Error calculating unique path count: %w", error)
	}

	return uniquePathCount, nil
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

func uniquePaths(matrix matrix, cache pathCountCache) (int, error) {
	rootNodeCoordinate, error := manifoldEntrace(matrix)

	if error != nil {
		return 0, error
	}

	return uniquePathsFromCoordinate(matrix, rootNodeCoordinate, cache), nil
}

func uniquePathsFromCoordinate(m matrix, c coordinate, cache pathCountCache) int {
	coordinateValue, found := m[c]

	if !found {
		return 1
	}

	switch coordinateValue {
	case start, empty:
		return uniquePathsFromCoordinate(m, c.south(), cache)
	case manifold:
		pathsToEast, found := cache[c.east()]
		if !found {
			pathsToEast = uniquePathsFromCoordinate(m, c.east(), cache)
			cache[c.east()] = pathsToEast
		}

		pathsToWest, found := cache[c.west()]
		if !found {
			pathsToWest = uniquePathsFromCoordinate(m, c.west(), cache)
			cache[c.west()] = pathsToWest
		}

		return pathsToEast + pathsToWest
	}

	panic(fmt.Sprintf("Unexpected coordinate value %q encountered at %v", coordinateValue, c))
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
