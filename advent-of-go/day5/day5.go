package day5

import (
	"cmp"
	"iter"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type freshIngredientRange struct {
	from, to ingedientId
}

func (f freshIngredientRange) all() iter.Seq[ingedientId] {
	return func(yield func(ingedientId) bool) {
		for i := f.from; i <= f.to; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func (f freshIngredientRange) contains(id ingedientId) bool {
	return f.from <= id && f.to >= id
}

func (f freshIngredientRange) len() int {
	return int(f.to) - int(f.from) + 1
}

func (this freshIngredientRange) overlapsWith(that freshIngredientRange) bool {
	ranges := []freshIngredientRange{this, that}
	slices.SortFunc(ranges, func(f1, f2 freshIngredientRange) int {
		return cmp.Compare(f1.from, f2.from)
	})

	return ranges[1].from <= ranges[0].to
}

func (this freshIngredientRange) merge(that freshIngredientRange) freshIngredientRange {
	return freshIngredientRange{min(this.from, that.from), max(this.to, that.to)}
}

func min(i1, i2 ingedientId) ingedientId {
	if i1 <= i2 {
		return i1
	}

	return i2
}

func max(i1, i2 ingedientId) ingedientId {
	if i1 >= i2 {
		return i1
	}

	return i2
}

type ingedientId int

func Run() (int, error) {

	bytes, error := os.ReadFile("day5/input.txt")
	if error != nil {
		return 0, error
	}

	freshIngredientRanges, _, error := parseInput(strings.Fields(string(bytes)))

	if error != nil {
		return 0, error
	}

	slices.SortFunc(freshIngredientRanges, func(a, b freshIngredientRange) int {
		return cmp.Compare(a.from, b.from)
	})

	mergedRanges := merge(freshIngredientRanges)

	lengths := lo.Map(mergedRanges, func(f freshIngredientRange, _ int) int {
		return f.len()
	})

	return lo.Sum(lengths), nil
}

func parseInput(lines []string) ([]freshIngredientRange, []ingedientId, error) {
	freshIngredientRanges := []freshIngredientRange{}
	ingredientIds := []ingedientId{}

	for _, line := range lines {
		maybeIngredientRange := strings.Split(line, "-")

		if len(maybeIngredientRange) == 2 {
			rangeFrom, error := strconv.Atoi(maybeIngredientRange[0])

			if error != nil {
				return nil, nil, error
			}

			rangeTo, error := strconv.Atoi(maybeIngredientRange[1])

			if error != nil {
				return nil, nil, error
			}

			freshIngredientRanges = append(freshIngredientRanges, freshIngredientRange{ingedientId(rangeFrom), ingedientId(rangeTo)})
			continue
		}

		ingredientId, error := strconv.Atoi(line)

		if error != nil {
			return nil, nil, error
		}

		ingredientIds = append(ingredientIds, ingedientId(ingredientId))
	}

	return freshIngredientRanges, ingredientIds, nil
}

func isFresh(freshIngredientsSorted []freshIngredientRange, ingredientId ingedientId) bool {
	for _, ingredientRange := range freshIngredientsSorted {
		if ingredientRange.from > ingredientId {
			return false
		}

		if ingredientRange.contains(ingredientId) {
			return true
		}
	}

	return false
}

func merge(ranges []freshIngredientRange) []freshIngredientRange {
	// assume ranges is sorted

	mergedRanges := []freshIngredientRange{}

	currentRange := ranges[0]

	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]
		if currentRange.overlapsWith(nextRange) {
			currentRange = currentRange.merge(nextRange)
		} else {
			mergedRanges = append(mergedRanges, currentRange)
			currentRange = nextRange
		}
	}

	mergedRanges = append(mergedRanges, currentRange)

	return mergedRanges
}
