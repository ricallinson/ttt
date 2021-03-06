package ttt

import (
	"errors"
	"regexp"
)

const (
	X = 'x'
	O = 'o'
)

type Game struct {
	moves      int
	board      []byte
	lastPlayer uint8
}

var win = map[uint8][]*regexp.Regexp{}

func init() {
	checks := map[uint8][]string{
		X: []string{
			"xxx......",
			"...xxx...",
			"......xxx",
			"x..x..x..",
			".x..x..x.",
			"..x..x..x",
			"x...x...x",
			"..x.x.x..",
		},
		O: []string{
			"ooo......",
			"...ooo...",
			"......ooo",
			"o..o..o..",
			".o..o..o.",
			"..o..o..o",
			"o...o...o",
			"..o.o.o..",
		},
	}
	for _, w := range checks[X] {
		regex, _ := regexp.Compile(w)
		win[X] = append(win[X], regex)
	}
	for _, w := range checks[O] {
		regex, _ := regexp.Compile(w)
		win[O] = append(win[O], regex)
	}
}

func CreateGame() *Game {
	return &Game{
		board: []byte("         "),
	}
}

// Place the given player at the provided position if it is available.
func (this *Game) Place(player uint8, pos int) (bool, error) {
	if pos < 0 || pos > 8 {
		return false, errors.New("Bad position")
	}
	if this.board[pos] != ' ' {
		return false, errors.New("Position taken")
	}
	if player != X && player != O {
		return false, errors.New("Not an 'x' or 'o'")
	}
	if player == this.lastPlayer {
		return false, errors.New("Wrong player")
	}
	this.moves++
	this.lastPlayer = player
	this.board[pos] = player
	for _, w := range win[player] {
		if w.Match(this.board) {
			return true, nil
		}
	}
	return false, nil
}

// Place the given player at the provided x, y if it is available.
func (this *Game) Move(player uint8, x int, y int) (bool, error) {
	pos := -1
	switch y {
	case 1:
		pos = x + y - 2
	case 2:
		pos = x + y
	case 3:
		pos = x + y + 2
	}
	return this.Place(player, pos)
}

func (this *Game) Board() []byte {
	return this.board
}

func (this *Game) Player() uint8 {
	if this.lastPlayer == X {
		return O
	}
	return X
}

func (this *Game) Draw() bool {
	if this.moves > 8 {
		return true
	}
	return false
}
