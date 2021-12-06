package day06

import (
	"fmt"
	"strings"

	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
)

func parseFile(lines []string) []int {
	var fishList []int

	for _, fish := range strings.Split(lines[0], ",") {
		fishList = append(fishList, util.MustAtoi(fish))
	}

	return fishList
}

func simulateFish(fishList []int, days int) int {
	for i := 0; i < days; i++ {
		tempList := fishList
		for i, fish := range tempList {
			if fish == 0 {
				tempList[i] = 6
				tempList = append(tempList, 8)
			} else {
				tempList[i]--
			}
		}
		fishList = tempList
		fmt.Println(i)
	}

	return len(fishList)
}
