package day04

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	lines, err := reader.ReadLines("day04/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	draw, boards := parseFile(lines)
	result := checkLastWinner(draw, boards)

	fmt.Printf("day04 part2 result: %d\n", result)
}
