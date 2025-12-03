package day3

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type battery struct {
	index   int
	joltage int
}
type batteryBank []battery

func Run() (int, error) {
	bytes, error := os.ReadFile("day3/input.txt")

	if error != nil {
		return 0, error
	}

	batteryBanks, error := parseInput(strings.Fields(string(bytes)))

	if error != nil {
		return 0, error
	}

	maxVoltages := lo.Map(batteryBanks, func(bBank batteryBank, _ int) int {
		return maxJoltage(bBank)
	})

	fmt.Println(maxVoltages)

	answer := lo.Sum(maxVoltages)

	return answer, nil
}

func parseInput(strings []string) ([]batteryBank, error) {
	batteryBanks := []batteryBank{}
	for _, str := range strings {
		batteryBank := batteryBank{}
		for index, rune := range str {
			joltage, error := strconv.Atoi(string(rune))

			if error != nil {
				return nil, error
			}

			batteryBank = append(batteryBank, battery{index, joltage})
		}
		batteryBanks = append(batteryBanks, batteryBank)
	}

	return batteryBanks, nil
}

func maxJoltage(bBank batteryBank) int {

	allButFinalBattery := bBank[:len(bBank)-1]

	tensDigitBattery := slices.MaxFunc(allButFinalBattery, func(b1, b2 battery) int {
		return cmp.Compare(b1.joltage, b2.joltage)
	})

	onesDigitCandidates := bBank[tensDigitBattery.index+1:]

	onesDigitBattery := slices.MaxFunc(onesDigitCandidates, func(b1, b2 battery) int {
		return cmp.Compare(b1.joltage, b2.joltage)
	})

	return (tensDigitBattery.joltage * 10) + onesDigitBattery.joltage
}
