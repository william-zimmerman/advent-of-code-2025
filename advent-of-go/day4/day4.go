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

	totalRemovedRolls := 0
	for {
		removedRolls := removeAccessibleRollsOfPaper(puzzleMap)
		totalRemovedRolls += removedRolls
		if removedRolls == 0 {
			break
		}
	}

	return totalRemovedRolls, nil
}

func removeAccessibleRollsOfPaper(puzzleMap puzzleMap) int {

	isPaper := predicate[coordinate](func(c coordinate) bool {
		return puzzleMap.isPaper(c)
	})

	isAccessible := predicate[coordinate](func(c coordinate) bool {
		paperNeighborCount := len(slices.Collect(filter(c.neighbors(), isPaper)))
		return paperNeighborCount < 4
	})

	accessibleRollsOfPaper := slices.Collect(filter(maps.Keys(puzzleMap), and(isPaper, isAccessible)))

	for _, coordinate := range accessibleRollsOfPaper {
		puzzleMap[coordinate] = empty
	}

	return len(accessibleRollsOfPaper)
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

func and[T any](p1, p2 predicate[T]) predicate[T] {
	return func(t T) bool {
		return p1(t) && p2(t)
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
