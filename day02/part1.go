package day02

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	lines, err := reader.ReadLines("day02/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	result := calcPos(lines)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("day02 part1 result: %d\n", result)
}
