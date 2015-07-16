package ttt

import (
	// "log"
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
func (this *Game) Place(player uint8, pos int) (bool, bool) {
	if this.board[pos] != ' ' {
		// log.Println("Position taken", pos)
		return false, false
	}
	if player != 'x' && player != 'o' {
		// log.Println("Not an 'x' or 'o'")
		return false, false
	}
	if player == this.lastPlayer {
		// log.Println("Wrong player")
		return false, false
	}
	this.lastPlayer = player
	this.board[pos] = player
	for _, w := range wins[player] {
		ok, _ := regexp.Match(w, this.board)
		if ok {
			return true, true
		}
	}
	return true, false
}

// Place the given player at the provided x, y if it is available.
func (this *Game) Move(player uint8, x int, y int) (bool, bool) {
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
		// log.Println("Bad position", pos)
		return false, false
	}
	return this.Place(player, pos)
}

func (this *Game) Board() string {
	return string(this.board)
}
