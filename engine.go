package main

import (
	"strconv"
)

func GetBestMove(board [3][3]string, order_of_move string) int {
	draw_moves := make([]int, 0, 9)
	lose_moves := make([]int, 0, 9)
	for index_y, _ := range board {
		for index_x, value := range board[index_y] {
			_, err := strconv.Atoi(value)
			if err == nil {
				cell := board[index_y][index_x]
				board[index_y][index_x] = order_of_move
				move := buildTree(board, order_of_move)
				if move == order_of_move {
					return (index_y)*3 + (index_x + 1)
				} else if move == "-" {
					draw_moves = append(draw_moves, (index_y)*3+(index_x+1))
				} else {
					lose_moves = append(lose_moves, (index_y)*3+(index_x+1))
				}
				board[index_y][index_x] = cell
			}
		}
	}

	if len(draw_moves) > 0 {
		return draw_moves[0]
	} else {
		return lose_moves[0]
	}
}

func buildTree(board [3][3]string, last_char string) string {
	winner := getWinner(board)
	if winner != "-" {
		return winner
	} else if isOver(board) {
		return "-"
	}

	current_char := "X"
	if last_char == "X" {
		current_char = "O"
	}

	win_moves := make([]int, 0, 9)
	lose_moves := make([]int, 0, 9)
	draw_moves := make([]int, 0, 9)
	for index_y, _ := range board {
		for index_x, value := range board[index_y] {
			_, err := strconv.Atoi(value)
			if err == nil {
				cell := board[index_y][index_x]
				board[index_y][index_x] = current_char
				result := buildTree(board, current_char)
				if current_char == result {
					win_moves = append(win_moves, (index_y)*3+(index_x+1))
				} else if result == "-" {
					draw_moves = append(draw_moves, (index_y)*3+(index_x+1))
				} else {
					lose_moves = append(lose_moves, (index_y)*3+(index_x+1))
				}
				board[index_y][index_x] = cell
			}
		}
	}
	if len(win_moves) > 0 {
		return current_char
	} else if len(draw_moves) > 0 {
		return "-"
	}
	return last_char
}

func isOver(board [3][3]string) bool {
	for index_y, _ := range board {
		for _, value := range board[index_y] {
			_, err := strconv.Atoi(value)
			if err == nil {
				return false
			}
		}
	}
	return true
}

func getWinner(board [3][3]string) string {
	for i := 0; i < 3; i++ {
		if board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		} else if board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}

	if board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}

	if board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return "-"
}
