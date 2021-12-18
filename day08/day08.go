package day08

import (
	"fmt"
	"strings"

	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
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

func decodeOutput(signals, output [][]string) int {
	var total int
	var digits = make(map[int]string, 10)

	for n, line := range signals {
		for _, code := range line {
			switch len(code) {
			case 2:
				digits[1] = code
			case 3:
				digits[7] = code
			case 4:
				digits[4] = code
			case 7:
				digits[8] = code
			}
		}
		for _, code := range line {
			switch len(code) {
			case 5:
				if countInCommon(code, digits[4]) == 2 {
					digits[2] = code

				} else if countInCommon(code, digits[1]) == 2 {
					digits[3] = code

				} else {
					digits[5] = code
				}

			case 6:
				if countInCommon(code, digits[1]) == 1 {
					digits[6] = code

				} else if countInCommon(code, digits[4]) == 4 {
					digits[9] = code

				} else {
					digits[0] = code
				}
			}
		}
		total += addOutputValues(digits, output[n])
	}
	return total
}

func countInCommon(a, b string) int {
	count := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				count++
			}
		}
	}
	return count
}

func addOutputValues(digits map[int]string, output []string) int {
	numString := ""
	for _, code := range output {
		for num, decoded := range digits {
			common := countInCommon(decoded, code)
			if common == len(decoded) && common == len(code) {
				numString += fmt.Sprintf("%d", num)
			}
		}
	}
	return util.MustAtoi(numString)
}
