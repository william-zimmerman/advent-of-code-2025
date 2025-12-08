package day8

import (
	"cmp"
	"fmt"
	"iter"
	"maps"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
	"github.com/samber/lo"
	"github.com/samber/lo/it"
)

type junctionBox struct {
	x, y, z int
}

type junctionBoxDistance struct {
	box1, box2 junctionBox
	distance   float64
}

type circuit map[junctionBox]struct{}

func (c circuit) add(box junctionBox) {
	c[box] = struct{}{}
}

func (c circuit) len() int {
	return it.Length(maps.Keys(c))
}

func (c circuit) string() string {

	sortedKeys := slices.Collect(maps.Keys(c))
	slices.SortFunc(sortedKeys, func(a, b junctionBox) int {
		return cmp.Compare(a.x, b.x)
	})

	builder := strings.Builder{}
	builder.WriteString("{")

	for k := range sortedKeys {
		builder.WriteString(fmt.Sprintf("%v, ", k))
	}
	builder.WriteString("}")
	return builder.String()
}

func combine(c1, c2 circuit) {
	for box := range maps.Keys(c2) {
		c1.add(box)
	}

	for box := range maps.Keys(c1) {
		c2.add(box)
	}
}

type circuitMap map[junctionBox]circuit

func (c circuitMap) addConnection(box1, box2 junctionBox) {
	box1Circuit, box1Found := c[box1]
	box2Circuit, box2Found := c[box2]

	if !box1Found && !box2Found {
		newCircuit := circuit{}
		newCircuit.add(box1)
		newCircuit.add(box2)
		c[box1] = newCircuit
		c[box2] = newCircuit
		return
	}

	if !box1Found {
		box2Circuit.add(box1)
		c[box1] = box2Circuit
	}

	if !box2Found {
		box1Circuit.add(box2)
		c[box2] = box1Circuit
	}

	if box1Found && box2Found {
		combine(box1Circuit, box2Circuit)
	}
}

func Run() (int, error) {

	bytes, err := os.ReadFile("day8/input.txt")

	if err != nil {
		return 0, fmt.Errorf("day8: failed to read input: %w", err)
	}

	junctionBoxes, err := parseInput(slices.Collect(strings.Lines(string(bytes))))
	if err != nil {
		return 0, fmt.Errorf("day8: failed to parse input: %w", err)
	}

	var distanceComparator utils.Comparator = func(a, b interface{}) int {
		return utils.Float64Comparator(a.(junctionBoxDistance).distance, b.(junctionBoxDistance).distance)
	}

	minHeap := binaryheap.NewWith(distanceComparator)

	for pair := range uniquePairs(junctionBoxes) {
		box1, box2 := pair[0], pair[1]
		distance := distance(box1, box2)
		minHeap.Push(junctionBoxDistance{box1, box2, distance})
	}

	circuitMap := circuitMap{}
	for range 1000 {
		value, ok := minHeap.Pop()
		if !ok {
			continue
		}
		circuitMap.addConnection(value.(junctionBoxDistance).box1, value.(junctionBoxDistance).box2)
	}

	uniqueCircuits := slices.Collect(it.UniqBy(maps.Values(circuitMap), func(c circuit) string {
		return c.string()
	}))
	slices.SortFunc(uniqueCircuits, func(a, b circuit) int {
		return -cmp.Compare(a.len(), b.len())
	})

	answer := lo.Product(lo.Map(uniqueCircuits[:3], func(c circuit, _ int) int { return c.len() }))

	// 47804 is too low
	return answer, nil
}

func parseInput(lines []string) ([]junctionBox, error) {

	coordinates := []junctionBox{}

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

		coordinates = append(coordinates, junctionBox{x, y, z})
	}

	return coordinates, nil
}

func distance(c1, c2 junctionBox) float64 {
	xSquare := math.Pow(float64(c1.x-c2.x), 2)
	ySquare := math.Pow(float64(c1.y-c2.y), 2)
	zSquare := math.Pow(float64(c1.z-c2.z), 2)

	return math.Sqrt(xSquare + ySquare + zSquare)
}

func uniquePairs[T any](slice []T) iter.Seq[[2]T] {

	return func(yield func([2]T) bool) {
		for i := 0; i < len(slice); i++ {
			for j := i + 1; j < len(slice); j++ {
				if !yield([2]T{slice[i], slice[j]}) {
					return
				}
			}
		}
	}
}
