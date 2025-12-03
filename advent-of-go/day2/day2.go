package day2

import (
	"fmt"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
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

	allIds := lo.FlatMap(idRanges, func(idRange idRange, _ int) []int {
		return slices.Collect(idRange.all())
	})

	allInvalidIds := lo.Filter(allIds, func(id int, _ int) bool {
		return isInvalid(id)
	})

	answer := lo.Sum(allInvalidIds)

	return answer, nil
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

func isInvalid(id int) bool {
	stringValue := strconv.Itoa(id)

	for i := 1; i <= len(stringValue)/2; i++ {
		if len(stringValue)%i != 0 {
			continue
		}
		prefix := stringValue[:i]
		repeatedPrefix := strings.Repeat(prefix, len(stringValue)/i)

		if repeatedPrefix == stringValue {
			return true
		}
	}

	return false
}
