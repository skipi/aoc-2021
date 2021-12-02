package day1

import (
	"fmt"

	"kutryj.io/aoc21/pkg/utils"
)

func Run() {
	rawMeasurements := utils.LoadData("./inputs/day1.txt")

	measurements := make([]int, 0)

	for _, rawMeasurement := range rawMeasurements {
		measurements = append(measurements, utils.ParseInt(rawMeasurement))
	}

	fmt.Printf("\nPart 1: %d\n", Part1(measurements))

	fmt.Printf("Part 2: %d\n", Part2(measurements))
}

func Part1(measurements []int) int {
	return NumberOfIncreases(measurements)
}

func Part2(measurements []int) int {
	measurements = SplitMeasurements(measurements)

	return NumberOfIncreases(measurements)
}

func NumberOfIncreases(measurements []int) int {
	if len(measurements) < 2 {
		return 0
	}
	increasesCounter := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increasesCounter += 1
		}
	}

	return increasesCounter
}

func SplitMeasurements(measurements []int) []int {
	if len(measurements) < 4 {
		return []int{}
	}

	newMeasurements := make([]int, 0)

	for i := 2; i < len(measurements); i++ {
		newMeasurement := measurements[i] + measurements[i-1] + measurements[i-2]
		newMeasurements = append(newMeasurements, newMeasurement)
	}

	return newMeasurements
}
