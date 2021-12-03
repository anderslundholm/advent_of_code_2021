package day03

import (
	"log"
	"strconv"
)

func powerConsumption(lines []string) int64 {
	gamma, epsilon := mostLeastCommonBits(lines, '1')
	return binaryToBaseTwo(gamma) * binaryToBaseTwo(epsilon)
}

func oxygenGenRating(lines []string) int64 {
	oxygen := calcRating(lines, '1')
	co2 := calcRating(lines, '0')
	return binaryToBaseTwo(oxygen) * binaryToBaseTwo(co2)
}

func calcRating(lines []string, xbit rune) string {
	var oxygen string
	var co2 string
	var tmp string
	var ret string
	count := 0
	for len(lines) > 1 {
		var keep []string
		oxygen, co2, tmp = oxygenCO2Bits(lines, xbit)
		for _, line := range lines {
			if tmp[count] == line[count] {
				keep = append(keep, line)
			}
		}
		lines = keep
		count++
	}
	if xbit == '1' {
		ret = oxygen
	} else {
		ret = co2
	}
	return ret
}

func mostLeastCommonBits(lines []string, xbit rune) (string, string) {
	mostCommonBits := make(map[int]int)
	var most string
	var least string
	for _, line := range lines {
		for i, x := range line {
			if x == xbit {
				mostCommonBits[i] = mostCommonBits[i] + 10
			} else if mostCommonBits[i] == 0 {
				mostCommonBits[i] = 0
			}
		}
	}
	for i := 0; i < len(mostCommonBits); i++ {
		if mostCommonBits[i] < len(lines)*5 {
			most = most + "0"
			least = least + "1"
		} else if mostCommonBits[i] > len(lines)*5 {
			most = most + "1"
			least = least + "0"
		} else if xbit == '0' {
			most = most + "1"
			least = least + "0"
		} else if xbit == '1' {
			most = most + "1"
			least = least + "0"
		}
	}
	return most, least
}

func oxygenCO2Bits(lines []string, xbit rune) (string, string, string) {
	mostCommonBits := make(map[int]int)
	var most string
	var tmp string
	var least string
	for _, line := range lines {
		for i, x := range line {
			if x == '1' {
				mostCommonBits[i] = mostCommonBits[i] + 10
			} else if mostCommonBits[i] == 0 {
				mostCommonBits[i] = 0
			}
		}
	}
	for i := 0; i < len(mostCommonBits); i++ {
		if mostCommonBits[i] < len(lines)*5 {
			most = most + "0"
		} else if mostCommonBits[i] > len(lines)*5 {
			most = most + "1"
		} else {
			most = most + "1"
		}
	}
	for i := 0; i < len(mostCommonBits); i++ {
		if mostCommonBits[i] > len(lines)*5 {
			tmp = tmp + "0"
			least = least + "1"
		} else if mostCommonBits[i] < len(lines)*5 {
			tmp = tmp + "1"
			least = least + "0"
		} else {
			tmp = tmp + "1"
			least = least + "0"
		}
	}
	if xbit == '1' {
		tmp = most
	}
	return most, least, tmp
}

func binaryToBaseTwo(binary string) int64 {
	baseTwo, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert binary: %v\n", err)
	}
	return baseTwo
}
