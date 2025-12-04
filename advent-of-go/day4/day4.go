package day4

import (
	"iter"
	"maps"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	x, y int
}

type puzzleMap map[coordinate]string

func (p puzzleMap) isPaper(key coordinate) bool {
	v, found := p[key]

	return found && (v == paper)
}

const (
	empty = "."
	paper = "@"
)

func Run() (int, error) {

	bytes, error := os.ReadFile("day4/input.txt")

	if error != nil {
		return 0, error
	}

	puzzleMap := parseInput(strings.Fields(string(bytes)))

	isPaper := func(c coordinate) bool {
		return puzzleMap.isPaper(c)
	}

	accessibleRollsOfPaper := 0
	for coordinate := range filter(maps.Keys(puzzleMap), isPaper) {
		paperNeighborCount := len(slices.Collect(filter(coordinate.neighbors(), isPaper)))
		if paperNeighborCount < 4 {
			accessibleRollsOfPaper++
		}

	}

	return accessibleRollsOfPaper, nil
}

func parseInput(lines []string) puzzleMap {

	puzzleMap := puzzleMap{}

	for y, line := range lines {
		for x, char := range line {
			c := coordinate{x: x, y: y}
			puzzleMap[c] = string(char)
		}
	}

	return puzzleMap
}

func (c coordinate) neighbors() iter.Seq[coordinate] {
	return func(yield func(coordinate) bool) {
		for _, neighbor := range []coordinate{c.north(), c.north().east(), c.east(), c.south().east(), c.south(), c.south().west(), c.west(), c.north().west()} {
			if !yield(neighbor) {
				return
			}
		}
	}
}

type predicate[T any] func(t T) bool

func filter[T any](items iter.Seq[T], p predicate[T]) iter.Seq[T] {
	return func(yield func(t T) bool) {
		for item := range items {
			if p(item) && !yield(item) {
				return
			}
		}
	}
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
