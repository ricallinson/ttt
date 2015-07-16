package ttt

import (
	"errors"
	"regexp"
)

type Game struct {
	board      []byte
	lastPlayer uint8
}

var wins = map[uint8][]string{
	'x': []string{
		"xxx......",
		"...xxx...",
		"......xxx",
		"x..x..x..",
		".x..x..x.",
		"..x..x..x",
		"x...x...x",
		"..x.x.x..",
	},
	'o': []string{
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

func CreateGame() *Game {
	return &Game{
		board: []byte("         "),
	}
}

// Place the given player at the provided position if it is available.
func (this *Game) Place(player uint8, pos int) (bool, error) {
	if this.board[pos] != ' ' {
		return false, errors.New("Position taken")
	}
	if player != 'x' && player != 'o' {
		return false, errors.New("Not an 'x' or 'o'")
	}
	if player == this.lastPlayer {
		return false, errors.New("Wrong player")
	}
	this.lastPlayer = player
	this.board[pos] = player
	for _, w := range wins[player] {
		ok, _ := regexp.Match(w, this.board)
		if ok {
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
	// log.Printf("Player: %s, x: %d, y: %d, pos: %d\n", string(player), x, y, pos)
	if pos < 0 || pos > 8 {
		return false, errors.New("Bad position")
	}
	return this.Place(player, pos)
}

func (this *Game) Board() string {
	return string(this.board)
}

func (this *Game) Player() uint8 {
	if this.lastPlayer == 'x' {
		return 'o'
	}
	return 'x'
}
