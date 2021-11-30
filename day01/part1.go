package day01

import (
	"fmt"
	"log"

	"github.com/anderslundholm/advent_of_code_2021/pkg/reader"
	"github.com/anderslundholm/advent_of_code_2021/pkg/timer"
)

// Part1 ...
func Part1() {
	defer timer.ExecutionTimer("Part1")()

	ints, err := reader.ReadInts("day01/input.txt")
	if err != nil {
		log.Fatalf("Could not read ints: %v\n", err)
	}

	product, err := partOne(ints)
	if err != nil {
		log.Fatalf("Could not find value: %v\n", err)
	}

	fmt.Printf("Day01 Part1 result: %d\n", product)
}
