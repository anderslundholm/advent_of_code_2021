package day01

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	ints, err := reader.ReadInts("day01/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	count := partTwo(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day01 Part2 result: %d\n", count)
}
