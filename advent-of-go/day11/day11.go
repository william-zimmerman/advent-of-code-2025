package day11

import (
	"fmt"
	"os"
	"strings"

	"github.com/samber/lo"
)

const (
	you                      = device("you")
	out                      = device("out")
	server                   = device("svr")
	digitalAnalogueConverter = device("dac")
	fastFourierTransform     = device("fft")
)

type device string

func Run() (int, error) {

	data, err := os.ReadFile("day11/input.txt")
	if err != nil {
		return 0, fmt.Errorf("error read input: %w", err)
	}

	deviceMap, err := parseInput(string(data))

	if err != nil {
		return 0, fmt.Errorf("error parsing input: %w", err)
	}

	devicePaths := pathsToOut(server, deviceMap)
	return devicePaths, nil
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

	type cacheKey struct {
		device   device
		foundDac bool
		foundFft bool
	}

	cache := make(map[cacheKey]int, len(deviceMap)*4)

	var pathsToOutRecurse func(d device, foundDac bool, foundFft bool) int
	pathsToOutRecurse = func(d device, foundDac bool, foundFft bool) int {

		fmt.Println(d)

		if d == out {
			if foundDac && foundFft {
				return 1
			} else {
				return 0
			}
		}

		key := cacheKey{d, foundDac, foundFft}
		if cachedValue, cacheHit := cache[key]; cacheHit {
			return cachedValue
		}

		childDevices, found := deviceMap[d]
		if !found {
			return 0
		}

		isFft := d == fastFourierTransform
		isDac := d == digitalAnalogueConverter

		sumOfChildren := 0
		for _, child := range childDevices {
			sumOfChildren += pathsToOutRecurse(child, foundDac || isDac, foundFft || isFft)
		}

		cache[key] = sumOfChildren

		return sumOfChildren
	}

	return pathsToOutRecurse(start, false, false)
}
