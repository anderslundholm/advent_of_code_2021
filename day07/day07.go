package day07

import (
	"math"
	"sort"
)

func calcPos(pos []int) int {
	var sum int
	sort.Ints(pos)
	mid := pos[len(pos)/2]

	for _, x := range pos {
		sum += int(math.Abs(float64(x - mid)))
	}
	return sum
}

func calcPos2(pos []int) int {
	sort.Ints(pos)
	largest := pos[len(pos)-1]
	smallest := math.MaxInt32
	var allSums []int

	for x := 0; x <= largest; x++ {
		sum := 0
		for _, y := range pos {
			n := int(math.Abs(float64(x - y)))
			sum += (n * (n + 1)) / 2
		}
		allSums = append(allSums, sum)
		if sum < smallest {
			smallest = sum
		}
	}
	return smallest
}

func expCost(move int) int {
	return move * (move + 1) / 2
}
