package day08

import (
	"strings"
)

func parseInput(lines []string) ([][]string, [][]string) {
	var signals, output [][]string
	for _, line := range lines {
		subLine := strings.Split(line, " | ")
		s := strings.Split(subLine[0], " ")
		o := strings.Split(subLine[1], " ")
		signals = append(signals, s)
		output = append(output, o)
	}
	return signals, output
}

func countUniqueDigits(output [][]string) int {
	count := 0
	for _, x := range output {
		for _, y := range x {
			switch len(y) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return count
}
