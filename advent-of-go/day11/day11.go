package day11

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/samber/lo"
	"github.com/samber/lo/it"
)

const (
	you = device("you")
	out = device("out")
)

type device string

type set[T comparable] map[T]struct{}

type devicePath string

func (d devicePath) add(tail device) devicePath {
	return devicePath(fmt.Sprintf("%s;%s", d, tail))
}

func (d devicePath) contains(target device) bool {
	return strings.Contains(string(d), string(target))
}

func newPath() devicePath {
	return ""
}

func Run() (int, error) {

	data, err := os.ReadFile("day11/input.txt")
	if err != nil {
		return 0, fmt.Errorf("error read input: %w", err)
	}

	deviceMap, err := parseInput(string(data))

	if err != nil {
		return 0, fmt.Errorf("error parsing input: %w", err)
	}

	answer := pathsToOut(you, deviceMap)

	return answer, nil
}

func parseInput(fileContents string) (map[device][]device, error) {

	deviceMap := make(map[device][]device)

	parseLine := func(line string) error {
		headAndTail := strings.Split(line, ":")

		if len(headAndTail) != 2 {
			return fmt.Errorf("expecting one colon in line %q", line)
		}

		sourceDevice := headAndTail[0]
		destinationDevices := lo.Map(strings.Fields(headAndTail[1]), func(item string, _ int) device {
			return device(item)
		})

		deviceMap[device(sourceDevice)] = destinationDevices

		return nil
	}

	lines := strings.Lines(fileContents)

	for line := range lines {
		err := parseLine(line)
		if err != nil {
			return nil, fmt.Errorf("error parsing: %w", err)
		}
	}

	return deviceMap, nil
}

func pathsToOut(start device, deviceMap map[device][]device) int {

	var pathsToOutRecurse func(start device, circuitMap map[device][]device, currentPath devicePath) int
	pathsToOutRecurse = func(start device, circuitMap map[device][]device, currentPath devicePath) int {

		if start == out {
			return 1
		}

		if currentPath.contains(start) {
			return 0
		}

		connectedDevices, found := circuitMap[start]

		if !found {
			return 0
		}

		paths := it.Sum(it.Map(slices.Values(connectedDevices), func(d device) int {
			return pathsToOutRecurse(d, circuitMap, currentPath.add(start))
		}))

		return paths
	}

	return pathsToOutRecurse(start, deviceMap, newPath())
}
