package day06

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part2 ...
func Part2() {
	defer timer.ExecutionTimer("Part2")()

	fishList, err := reader.ReadCommaSeparatedInts("day06/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	result := simulateFish2(fishList, 256)

	fmt.Printf("day06 part2 result: %d\n", result)
}
