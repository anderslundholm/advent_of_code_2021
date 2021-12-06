package day06

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day06/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	fishList := parseFile(lines)
	result := simulateFish2(fishList, 80)

	fmt.Printf("day06 part1 result: %d\n", result)
}
