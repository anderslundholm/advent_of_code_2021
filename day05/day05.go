package day05

import (
	"errors"
	"strings"

	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
)

type coord struct {
	x int
	y int
}

func parseFile(lines []string, diag bool) []coord {
	var allCoords []coord
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		c1 := strings.Split(coords[0], ",")
		c2 := strings.Split(coords[1], ",")
		start := coord{x: util.MustAtoi(c1[0]), y: util.MustAtoi(c1[1])}
		end := coord{x: util.MustAtoi(c2[0]), y: util.MustAtoi(c2[1])}
		coordRange, err := calcCoordRange(start, end, diag)
		if err != nil {
			continue
		}
		allCoords = append(allCoords, coordRange...)
	}
	return allCoords
}

func calcOverlap(allCords []coord) int {
	count := 0
	var overlap = make(map[coord]int)
	for _, c := range allCords {
		overlap[c] += 1
	}
	for _, point := range overlap {
		if point > 1 {
			count++
		}
	}
	return count
}

func calcCoordRange(start coord, end coord, diag bool) ([]coord, error) {
	var coordRange []coord
	x := true
	if !diag && (start.x != end.x && start.y != end.y) {
		err := errors.New("diagonal line")
		return []coord{}, err
	} else if start.x != end.x && start.y != end.y {
		dx, dy := util.Sgn(float64(end.x-start.x)), util.Sgn(float64(end.y-start.y))
		x, y := start.x, start.y
		for x != end.x+dx {
			coordRange = append(coordRange, coord{x: x, y: y})
			x += dx
			y += dy
		}
	} else {
		from := start.x
		to := end.x
		if start.x == end.x {
			from = start.y
			to = end.y
			x = false
		}
		if from > to {
			from, to = to, from
		}
		for i := 0; i <= to-from; i++ {
			if x {
				coordRange = append(coordRange, coord{x: from + i, y: start.y})
			} else {
				coordRange = append(coordRange, coord{x: start.x, y: from + i})
			}
		}
	}
	return coordRange, nil
}
