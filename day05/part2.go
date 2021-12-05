package day05

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day05/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	allCoords := parseFile(lines, true)
	result := calcOverlap(allCoords)

	fmt.Printf("day05 part2 result: %d\n", result)
}
