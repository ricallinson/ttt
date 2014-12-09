package main
import (
	"fmt"
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
		}
		if count == 3 {
			return player
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

func renderBoard(board []int) {
	row := 0
	for _, pos := range board {
		if row == 2 {
			row = 0
			fmt.Println(renderPosition(pos))
		} else {
			row++
			fmt.Print(renderPosition(pos))
		}
	}
}

func makeMove(value int, pos int, board []int) bool {
	if pos >= len(board) || board[pos] > 0 {
		return false
	}
	board[pos] = value
	return true
}

func loop(board []int, wins [][]int) {
	turn := false
	move := 0
	winner := 0
	for {
		renderBoard(board)
		winner = testForWin(board, wins)
		if winner == 1 {
			fmt.Println("O Wins!")
			return
		}
		if winner == 2 {
			fmt.Println("X Wins!")
			return
		}
		if turn == false {
			for turn == false {
				fmt.Print("O: ")
				fmt.Scanf("%d", &move)
				if makeMove(1, move, board) {
					turn = true
				}
			}
		} else {
			for turn {
				fmt.Print("X: ")
				fmt.Scanf("%d", &move)
				if makeMove(2, move, board) {
					turn = false
				}
			}
		}
	}
}

func main() {
	loop(getBoard(), getWins())
}
