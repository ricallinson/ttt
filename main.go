package main

import (
	"fmt"
	"os"
)

// 012
// 345
// 678
func getBoard() []int {
	return []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
}

func getWins() [][]int {
	return [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
		[]int{0, 3, 6},
		[]int{1, 2, 7},
		[]int{2, 5, 8},
		[]int{0, 4, 8},
		[]int{2, 4, 6},
	}
}

func canMove(board []int) bool {
	count := 0
	for _, v := range board {
		if v > 0 {
			count++
		}
	}
	if count == 9 {
		return false
	}
	return true
}

// 0 == No win
// 1 == O Wins
// 2 == X Wins
func testForWin(board []int, wins [][]int) int {
	player := 0
	count := 0
	for _, win := range wins {
		for _, pos := range win {
			if player == 0 {
				player = board[pos]
			}
			if player > 0 && player == board[pos] {
				count++
			}
			if count == 3 {
				return player
			}
		}
		player = 0
		count = 0
	}
	return player
}

func renderPosition(pos int) string {
	if pos == 1 {
		return "O"
	}
	if pos == 2 {
		return "X"
	}
	return " "
}

func renderBoard(board []int) string {
	screen := ""
	row := 0
	for _, pos := range board {
		screen = screen + renderPosition(pos)
		row++
		if row == 3 {
			row = 0
			screen = screen + "\n"
		}
	}
	return screen
}

func makeMove(value int, pos int, board []int) bool {
	if value < 1 || value > 2 {
		return false
	}
	if pos >= len(board) || board[pos] > 0 {
		return false
	}
	board[pos] = value
	return true
}

func loop(input *os.File, output *os.File, board []int, wins [][]int) {
	turn := false
	move := 0
	winner := 0
	for {
		fmt.Fprintln(output, renderBoard(board))
		winner = testForWin(board, wins)
		if winner == 1 {
			fmt.Fprintln(output, "O Wins!")
			return
		}
		if winner == 2 {
			fmt.Fprintln(output, "X Wins!")
			return
		}
		if canMove(board) == false {
			fmt.Fprintln(output, "Draw!")
			return
		}
		if turn == false { // O's turn
			for turn == false {
				fmt.Fprint(output, "O: ")
				fmt.Fscanf(input, "%d", &move)
				if makeMove(1, move, board) {
					turn = true
				}
			}
		} else { // X's turn
			for turn {
				fmt.Fprint(output, "X: ")
				fmt.Fscanf(input, "%d", &move)
				if makeMove(2, move, board) {
					turn = false
				}
			}
		}
	}
}

func main() {
	loop(os.Stdin, os.Stdout, getBoard(), getWins())
}
