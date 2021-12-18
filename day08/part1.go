package day08

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day08/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}
	_, output := parseInput(lines)
	result := countUniqueDigits(output)

	fmt.Printf("day08 part1 result: %d\n", result)
}
