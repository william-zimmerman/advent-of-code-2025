package day9

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo/it"
)

type coordinate struct {
	x, y int
}

func Run() (int, error) {
	fmt.Println("Hello from day9")

	bytes, fileReadErr := os.ReadFile("day9/input.txt")

	if fileReadErr != nil {
		return 0, fileReadErr
	}

	coordiantes, parseErr := parseInput(bytes)

	if parseErr != nil {
		return 0, parseErr
	}

	answer := it.Max(it.Map(uniquePairs(coordiantes), areaA))

	return answer, nil
}

func parseInput(fileContents []byte) ([]coordinate, error) {

	parseLine := func(line string) (coordinate, error) {
		vals := strings.Split(strings.TrimRight(line, "\n"), ",")
		if len(vals) != 2 {
			return coordinate{}, fmt.Errorf("Expecting two values separated by comma; found %q", line)
		}

		x, xErr := strconv.Atoi(vals[0])

		if xErr != nil {
			return coordinate{}, fmt.Errorf("Error parsing x value of %q : %w", line, xErr)
		}

		y, yErr := strconv.Atoi(vals[1])

		if yErr != nil {
			return coordinate{}, fmt.Errorf("Error parsing y value of %q : %w", line, yErr)
		}

		return coordinate{x, y}, nil
	}

	coordinates := []coordinate{}

	for line := range strings.Lines(string(fileContents)) {
		coordinate, err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("Error parsing lines: %w", err)
		}

		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func area(c1, c2 coordinate) int {
	return (abs(c1.x-c2.x) + 1) * (abs(c1.y-c2.y) + 1)
}

func areaA(pair [2]coordinate) int {
	return area(pair[0], pair[1])
}

func uniquePairs[T any](slice []T) iter.Seq[[2]T] {
	return func(yield func([2]T) bool) {
		for i := range slice {
			for j := i + 1; j < len(slice); j++ {
				if !yield([2]T{slice[i], slice[j]}) {
					return
				}
			}
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
