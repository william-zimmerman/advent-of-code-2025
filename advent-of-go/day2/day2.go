package day2

import (
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

type idRange struct {
	first int
	last  int
}

func (ir idRange) all() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := ir.first; i <= ir.last; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Run() (int, error) {

	bytes, error := os.ReadFile("day2/input.txt")

	if error != nil {
		return 0, error
	}

	idRanges, error := parseInput(strings.TrimSuffix(string(bytes), "\n"))

	if error != nil {
		return 0, error
	}

	allInvalidIds := []int{}
	for _, idRange := range idRanges {
		allInvalidIds = append(allInvalidIds, invalidIds(idRange)...)
	}

	sumOfInvalidIds := 0
	for _, id := range allInvalidIds {
		sumOfInvalidIds += id
	}

	return sumOfInvalidIds, nil
}

func parseInput(input string) ([]idRange, error) {

	parseRange := func(stringRange string) (idRange, error) {
		ids := strings.Split(stringRange, "-")

		if len(ids) != 2 {
			return idRange{}, fmt.Errorf("error parsing range string %q; expecting two ids separated by hyphen", stringRange)
		}

		first, error := strconv.Atoi(ids[0])

		if error != nil {
			return idRange{}, error
		}

		last, error := strconv.Atoi(ids[1])

		if error != nil {
			return idRange{}, error
		}

		return idRange{first, last}, nil
	}

	stringRanges := strings.Split(input, ",")

	var idRanges []idRange = make([]idRange, 0, len(stringRanges))
	for _, stringRange := range stringRanges {
		idRange, error := parseRange(stringRange)

		if error != nil {
			return nil, error
		}

		idRanges = append(idRanges, idRange)
	}

	return idRanges, nil
}

func invalidIds(idRange idRange) []int {

	invalidIds := []int{}

	for id := range idRange.all() {
		stringValue := strconv.Itoa(id)

		if len(stringValue)%2 != 0 {
			continue
		}

		midpoint := len(stringValue) / 2

		firstHalf := stringValue[:midpoint]
		secondHalf := stringValue[midpoint:]

		if firstHalf == secondHalf {
			invalidIds = append(invalidIds, id)
		}
	}

	return invalidIds
}
