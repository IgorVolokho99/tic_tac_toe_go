package main

import (
	"fmt"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
)

func main() {
	mapField := [3][3]string{
		{"1", "X", "3"},
		{"4", "5", "O"},
		{"7", "8", "9"},
	}

	printMap(mapField)
}

func printMap(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch board[i][j] {
			case "X":
				fmt.Print(Red, "X", Reset)
			case "O":
				fmt.Print(Blue, "O", Reset)
			default:
				fmt.Print(board[i][j])
			}

			if j != 2 {
				fmt.Print(" ")
			} else {
				fmt.Print("\n")
			}
		}
	}
}

func isMoveCorrect(board [3][3]string, move string) bool {
	movei := byte(move[0]) - 48
	cell := board[(movei-1)/3][(movei-1)%3]
	if cell == "X" || cell == "O" {
		return false
	} else {
		return true
	}
}

func getPlayerCellNumber(board [3][3]string) string {
	var move string

	for {
		fmt.Scan(&move)

		if len(move) != 1 || move[0] < '1' || move[0] > '9' {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру от 1 до 9.")
			continue
		}

		if !isMoveCorrect(board, move) {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру пустой ячейки.")
			continue
		}

		return move
	}
}
