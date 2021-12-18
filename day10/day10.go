package day10

import (
	"log"
	"sort"
)

func calcSyntaxErrors(lines [][]string) int {
	charMap := map[string]int{"(": 3, "[": 57, "{": 1197, "<": 25137, ")": 3, "]": 57, "}": 1197, ">": 25137}
	count := 0
Line:
	for _, line := range lines {
		var openChars []string
		for _, char := range line {
			switch char {
			case "(", "[", "{", "<":
				openChars = append(openChars, char)
			case ")", "]", "}", ">":
				if charMap[openChars[len(openChars)-1]] == charMap[char] {
					openChars = openChars[:len(openChars)-1]
				} else {
					count = count + charMap[char]
					continue Line
				}
			default:
				log.Fatal("Wrong input")
			}
		}
	}
	return count
}

func calcSyntaxErrors2(lines [][]string) int {
	charMap := map[string]int{"(": 1, "[": 2, "{": 3, "<": 4, ")": 1, "]": 2, "}": 3, ">": 4}
	// endChars := map[int]string{1: ")", 2: "]", 3: "}", 4: ">"}
	var totalScores []int
	var finalScore int
Line:
	for _, line := range lines {
		var openChars []string
		for _, char := range line {
			switch char {
			case "(", "[", "{", "<":
				openChars = append(openChars, char)
			case ")", "]", "}", ">":
				if charMap[openChars[len(openChars)-1]] == charMap[char] {
					openChars = openChars[:len(openChars)-1]
				} else {
					continue Line
				}
			default:
				log.Fatal("Wrong input")
			}
		}
		score := 0
		for i := len(openChars) - 1; i >= 0; i-- {
			score = score*5 + charMap[openChars[i]]
		}
		totalScores = append(totalScores, score)
	}
	sort.Ints(totalScores)
	finalScore = totalScores[len(totalScores)/2]
	return finalScore
}
