package day04

import (
	"errors"
	"strings"

	"github.com/anderslundholm/advent_of_code_2021/pkg/util"
)

func parseFile(lines []string) ([]string, [][]string) {
	var boards [][]string
	draw := strings.Split(lines[0], ",")
	for i := 2; i < len(lines); i += 6 {
		var board []string
		boardLine := strings.Join(lines[i:i+5], " ")
		for _, j := range strings.Split(boardLine, " ") {
			if j == "" {
				continue
			}
			board = append(board, j)
		}
		boards = append(boards, board)
	}
	return draw, boards
}

func checkBoards(draw []string, boards [][]string) (int, error) {
	err := errors.New("no value")
	for _, number := range draw {
		for _, board := range boards {
			for i, n := range board {
				if n == number {
					board[i] = "0"
				}
			}
			if checkWinner(board) {
				sum := 0
				for _, j := range board {
					sum += util.MustAtoi(j)
				}
				return sum * util.MustAtoi(number), nil
			}
		}
	}
	return 0, err
}

func checkLastWinner(draw []string, boards [][]string) int {
	var result int
	winners := make([]bool, len(boards))
	for _, number := range draw {
		for x, board := range boards {
			if winners[x] {
				continue
			}
			for i, n := range board {
				if n == number {
					boards[x][i] = "0"
					break
				}
			}
			if checkWinner(boards[x]) {
				sum := 0
				for _, j := range boards[x] {
					sum += util.MustAtoi(j)
				}
				winners[x] = true
				result = sum * util.MustAtoi(number)
			}
		}
	}
	return result
}

func checkWinner(board []string) bool {
	for i := 0; i < 5; i++ {
		win := true
		for j := i * 5; j < (i+1)*5; j++ {
			if board[j] != "0" {
				win = false
				break
			}
		}
		if win {
			return true
		}
		win = true
		for j := i; j < 25; j += 5 {
			if board[j] != "0" {
				win = false
				break
			}
		}
		if win {
			return true
		}
	}
	return false
}
