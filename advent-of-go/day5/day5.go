package day5

import (
	"cmp"
	"os"
	"slices"
	"strconv"
	"strings"
)

type freshIngredientRange struct {
	from, to ingedientId
}

func (f freshIngredientRange) contains(id ingedientId) bool {
	return f.from <= id && f.to >= id
}

type ingedientId int

func Run() (int, error) {

	bytes, error := os.ReadFile("day5/input.txt")
	if error != nil {
		return 0, error
	}

	freshIngredients, ingredientIds, error := parseInput(strings.Fields(string(bytes)))

	if error != nil {
		return 0, error
	}

	slices.SortFunc(freshIngredients, func(a, b freshIngredientRange) int {
		return cmp.Compare(a.from, b.from)
	})

	freshIngredientCount := 0
	for _, ingredientId := range ingredientIds {
		if isFresh(freshIngredients, ingredientId) {
			freshIngredientCount++
		}
	}

	return freshIngredientCount, nil
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
