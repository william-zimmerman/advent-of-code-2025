package day8

import (
	"advent-of-go/day8/types"
	"fmt"
	"iter"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

func Run() (int, error) {

	bytes, err := os.ReadFile("day8/input.txt")

	if err != nil {
		return 0, fmt.Errorf("day8: failed to read input: %w", err)
	}

	junctionBoxes, err := parseInput(slices.Collect(strings.Lines(string(bytes))))
	if err != nil {
		return 0, fmt.Errorf("day8: failed to parse input: %w", err)
	}

	var distanceComparator utils.Comparator = func(a, b any) int {
		return utils.Float64Comparator(a.(types.JunctionBoxDistance).Distance, b.(types.JunctionBoxDistance).Distance)
	}

	// TODO: replace this min heap with a slice and then sort
	junctionBoxDistanceMinHeap := binaryheap.NewWith(distanceComparator)

	for pair := range uniquePairs(junctionBoxes) {
		box1, box2 := pair[0], pair[1]
		distance := distance(box1, box2)
		junctionBoxDistanceMinHeap.Push(types.JunctionBoxDistance{Box1: box1, Box2: box2, Distance: distance})
	}

	circuitMap := types.InitCircuitMap(junctionBoxes...)

	lastJunctionBoxesToBeConnected := types.JunctionBoxDistance{}
	for circuitMap.UniqueCircuitCount() > 1 {
		untyped, poppedElement := junctionBoxDistanceMinHeap.Pop()

		if poppedElement {
			junctionBoxDistance := untyped.(types.JunctionBoxDistance)
			lastJunctionBoxesToBeConnected = junctionBoxDistance
			circuitMap.Connect(junctionBoxDistance.Box1, junctionBoxDistance.Box2)
		} else {
			return 0, fmt.Errorf("Ran out of junction boxes to connect")
		}
	}

	answer := lastJunctionBoxesToBeConnected.Box1.X * lastJunctionBoxesToBeConnected.Box2.X

	return answer, nil
}

func parseInput(lines []string) ([]types.JunctionBox, error) {

	coordinates := []types.JunctionBox{}

	for _, line := range lines {
		stringCoordinates := strings.Split(strings.TrimRight(line, "\n"), ",")
		if len(stringCoordinates) != 3 {
			return nil, fmt.Errorf("expected 3 stringCoordinates, got %d from input %q", len(stringCoordinates), line)
		}

		x, xErr := strconv.Atoi(stringCoordinates[0])
		if xErr != nil {
			return nil, fmt.Errorf("could not parse x junctionBox: %w", xErr)
		}

		y, yErr := strconv.Atoi(stringCoordinates[1])
		if yErr != nil {
			return nil, fmt.Errorf("could not parse y junctionBox: %w", yErr)
		}

		z, zErr := strconv.Atoi(stringCoordinates[2])
		if zErr != nil {
			return nil, fmt.Errorf("could not parse z junctionBox: %w", zErr)
		}

		coordinates = append(coordinates, types.JunctionBox{X: x, Y: y, Z: z})
	}

	return coordinates, nil
}

func distance(c1, c2 types.JunctionBox) float64 {
	xSquare := math.Pow(float64(c1.X-c2.X), 2)
	ySquare := math.Pow(float64(c1.Y-c2.Y), 2)
	zSquare := math.Pow(float64(c1.Z-c2.Z), 2)

	return math.Sqrt(xSquare + ySquare + zSquare)
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
