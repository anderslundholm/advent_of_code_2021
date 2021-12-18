package day09

import (
	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
)

type coord struct {
	x, y int
}

func findNeighbours(lines [][]int, x, y int) []int {
	var neightbours []int
	if x+1 < len(lines[0]) {
		neightbours = append(neightbours, lines[y][x+1])
	}
	if x-1 >= 0 {
		neightbours = append(neightbours, lines[y][x-1])
	}
	if y+1 < len(lines) {
		neightbours = append(neightbours, lines[y+1][x])
	}
	if y-1 >= 0 {
		neightbours = append(neightbours, lines[y-1][x])
	}
	return neightbours
}

func calcRiskLevels(lines [][]int) int {
	result := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			neightbours := findNeighbours(lines, x, y)
			min, _ := util.MinMax(neightbours)
			if lines[y][x] < min {
				result += 1 + lines[y][x]
			}
		}
	}
	return result
}
