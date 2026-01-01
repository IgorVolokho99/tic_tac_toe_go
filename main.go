package main

import (
	"fmt"
	"strconv"
)

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
)

func main() {
	board := [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	order_of_move := "X"
	for i := 0; i < 9; i++ {
		if order_of_move == "X" {
			fmt.Println("Ходят крестики")
		} else {
			fmt.Println("Ходят нолики")
		}
		printMap(board)
		move := ""
		if order_of_move == "X" {
			move = getPlayerCellNumber(board)
		} else {
			move = strconv.Itoa(GetBestMove(board, "O"))
		}
		board = makeMove(board, move, order_of_move == "O")
		if hasWinner(board) {
			printMap(board)
			if order_of_move == "X" {
				fmt.Println("Крестики победили!")
			} else {
				fmt.Println("Нолики победили!")
			}
			break
		}
		if order_of_move == "X" {
			order_of_move = "O"
		} else {
			order_of_move = "X"
		}
	}
	if !hasWinner(board) {
		fmt.Println("Ничья!")
	}
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

func makeMove(board [3][3]string, cellNumber string, isZeroMow bool) [3][3]string {
	cellNumberi := byte(cellNumber[0]) - 48
	if isZeroMow {
		board[(cellNumberi-1)/3][(cellNumberi-1)%3] = "O"
	} else {
		board[(cellNumberi-1)/3][(cellNumberi-1)%3] = "X"
	}
	return board
}

func hasWinner(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return true
		} else if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return true
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return true
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return true
	}
	return false
}
