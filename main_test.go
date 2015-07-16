package ttt

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestGame(t *testing.T) {

	Describe("Game.Move()", func() {
		It("should return x as the winner", func() {
			g := CreateGame()
			g.Move(X, 1, 1)
			g.Move(O, 1, 3)
			g.Move(X, 2, 1)
			g.Move(O, 2, 3)
			win, _ := g.Move(X, 3, 1)
			AssertEqual(win, true)
		})
		It("should return o as the winner", func() {
			g := CreateGame()
			g.Move(O, 1, 1)
			g.Move(X, 1, 3)
			g.Move(O, 2, 1)
			g.Move(X, 2, 3)
			win, _ := g.Move(O, 3, 1)
			AssertEqual(win, true)
		})
		It("should return x as the winner again", func() {
			g := CreateGame()
			g.Move(X, 1, 1)
			g.Move(O, 1, 3)
			g.Move(X, 2, 2)
			g.Move(O, 2, 3)
			win, _ := g.Move(X, 3, 3)
			AssertEqual(win, true)
		})
		It("should return o as the winner again", func() {
			g := CreateGame()
			g.Move(O, 3, 1)
			g.Move(X, 2, 1)
			g.Move(O, 2, 2)
			g.Move(X, 2, 3)
			win, _ := g.Move(O, 1, 3)
			AssertEqual(win, true)
		})
		It("should return x as a false move", func() {
			g := CreateGame()
			g.Move(O, 1, 1)
			move, _ := g.Move(X, 1, 1)
			AssertEqual(move, false)
		})
		It("should return false as a bad actor", func() {
			g := CreateGame()
			move, _ := g.Move('q', 1, 2)
			AssertEqual(move, false)
		})
		It("should return false from low X cord", func() {
			g := CreateGame()
			move, _ := g.Move(X, 0, 1)
			AssertEqual(move, false)
		})
		It("should return false from low Y cord", func() {
			g := CreateGame()
			move, _ := g.Move(X, 1, 0)
			AssertEqual(move, false)
		})
		It("should return false high X cord", func() {
			g := CreateGame()
			move, _ := g.Move(X, 10, 1)
			AssertEqual(move, false)
		})
		It("should return false high Y cord", func() {
			g := CreateGame()
			move, _ := g.Move(X, 1, 10)
			AssertEqual(move, false)
		})
		It("should return false as same player is not allowed in a row", func() {
			g := CreateGame()
			g.Move(X, 1, 1)
			move, _ := g.Move(X, 1, 2)
			AssertEqual(move, false)
		})
	})

	Report(t)
}
