package day04

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day04/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	draw, boards := parseFile(lines)
	result, err := checkBoards(draw, boards)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("day04 part1 result: %d\n", result)
}
