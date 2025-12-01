package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type part1Answer = int
type direction int
type lockReading int

const (
	left direction = iota
	right
)

const (
	leftcode  string = "L"
	rightcode string = "R"
)

const safeDialNumberCount int = 100

type rotation struct {
	direction direction
	distance  int
}

func Run() (part1Answer, error) {
	bytes, error := os.ReadFile("day1/day1_input.txt")

	if error != nil {
		return 0, error
	}

	entireFileContents := string(bytes)
	lines := strings.Fields(entireFileContents)
	rotations, error := mapLines(lines)

	if error != nil {
		return 0, error
	}

	currentLockReading := lockReading(50)
	zeroReadingCount := 0
	for _, rotation := range rotations {
		currentLockReading = apply(rotation, currentLockReading)
		if currentLockReading == 0 {
			zeroReadingCount++
		}
	}

	return zeroReadingCount, nil
}

func mapLines(lines []string) ([]rotation, error) {
	rotations := []rotation{}
	for _, line := range lines {
		rotation, error := parseLine(line)
		if error != nil {
			return nil, error
		}
		rotations = append(rotations, rotation)
	}

	return rotations, nil
}

func parseLine(line string) (rotation, error) {
	directionCode := line[:1]
	distanceString := line[1:]
	distance, error := strconv.Atoi(distanceString)

	if error != nil {
		return rotation{}, error
	}

	direction, error := parseDirection(directionCode)

	if error != nil {
		return rotation{}, error
	}

	return rotation{direction, distance}, nil
}

func parseDirection(directionCode string) (direction, error) {
	switch directionCode {
	case leftcode:
		return left, nil
	case rightcode:
		return right, nil
	default:
		return -1, fmt.Errorf("Cannot parse direction code %q", directionCode)
	}
}

func apply(r rotation, lr lockReading) lockReading {
	switch r.direction {
	case right:
		return rotateRight(r.distance, lr)
	case left:
		return rotateLeft(r.distance, lr)
	default:
		panic(fmt.Sprintf("Unexpected rotation direction %q", r.direction))
	}
}

func rotateRight(distance int, currentReading lockReading) lockReading {
	newReading := (distance + int(currentReading)) % safeDialNumberCount
	return lockReading(newReading)
}

func rotateLeft(distance int, currentReading lockReading) lockReading {
	effectiveDistance := distance % safeDialNumberCount
	newReading := int(currentReading) - effectiveDistance
	if newReading < 0 {
		newReading += safeDialNumberCount
	}
	return lockReading(newReading)
}
