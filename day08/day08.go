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
	fmt.Printf("%s\n\n\n", signals)
	fmt.Println(output)
	var digits = make(map[int]string, 10)

	for _, line := range signals {
		for _, code := range line {
			switch len(code) {
			case 2:
				digits[1] = code
			case 3:
				digits[7] = code
			case 4:
				digits[4] = code
			case 5:

			case 6:

			case 7:
				digits[8] = code
			}
		}

	}

	fmt.Println(digits)
	return addOutputValues(digits, output)
}

func addOutputValues(digits map[int]string, output [][]string) int {
	result := 0

	for n, line := range output {
		number := ""
		for _, code := range line {
			for k := range digits {

				number += fmt.Sprintf("%d", k)

			}
			fmt.Println(n, code, number)

		}

		total := util.MustAtoi(number)
		result += total

	}

	return result
}
