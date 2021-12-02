package day02

import (
	"log"
	"strings"

	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
)

func calcPos(lines []string) int {
	calcCommands := map[string]int{"down": 0, "forward": 0}
	for _, x := range lines {
		commands := strings.Fields(x)
		value := commands[1]

		switch commands[0] {
		case "down":
			calcCommands["down"] = calcCommands["down"] + util.MustAtoi(value)
		case "up":
			calcCommands["down"] = calcCommands["down"] - util.MustAtoi(value)
		case "forward":
			calcCommands["forward"] = calcCommands["forward"] + util.MustAtoi(value)
		default:
			log.Fatalln("Unknown input.")
		}
	}
	return calcCommands["down"] * calcCommands["forward"]
}

func calcPosWithAim(lines []string) int {
	calcCommands := map[string]int{"down": 0, "forward": 0, "aim": 0}
	for _, x := range lines {
		commands := strings.Fields(x)
		value := commands[1]

		switch commands[0] {
		case "down":
			calcCommands["aim"] = calcCommands["aim"] + util.MustAtoi(value)
		case "up":
			calcCommands["aim"] = calcCommands["aim"] - util.MustAtoi(value)
		case "forward":
			calcCommands["forward"] = calcCommands["forward"] + util.MustAtoi(value)
			calcCommands["down"] = calcCommands["down"] + calcCommands["aim"]*util.MustAtoi(value)
		default:
			log.Fatalln("Unknown input.")
		}
	}
	return calcCommands["down"] * calcCommands["forward"]
}
