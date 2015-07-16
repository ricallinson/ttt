package ttt

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestGame(t *testing.T) {

	Describe("Game.Move()", func() {
		It("should return x as the winner", func() {
			g := CreateGame()
			g.Move('x', 1, 1)
			g.Move('o', 1, 3)
			g.Move('x', 2, 1)
			g.Move('o', 2, 3)
			win, _ := g.Move('x', 3, 1)
			AssertEqual(win, true)
		})
		It("should return o as the winner", func() {
			g := CreateGame()
			g.Move('o', 1, 1)
			g.Move('x', 1, 3)
			g.Move('o', 2, 1)
			g.Move('x', 2, 3)
			win, _ := g.Move('o', 3, 1)
			AssertEqual(win, true)
		})
		It("should return x as the winner again", func() {
			g := CreateGame()
			g.Move('x', 1, 1)
			g.Move('o', 1, 3)
			g.Move('x', 2, 2)
			g.Move('o', 2, 3)
			win, _ := g.Move('x', 3, 3)
			AssertEqual(win, true)
		})
		It("should return o as the winner again", func() {
			g := CreateGame()
			g.Move('o', 3, 1)
			g.Move('x', 2, 1)
			g.Move('o', 2, 2)
			g.Move('x', 2, 3)
			win, _ := g.Move('o', 1, 3)
			AssertEqual(win, true)
		})
		It("should return x as a false move", func() {
			g := CreateGame()
			g.Move('o', 1, 1)
			move, _ := g.Move('x', 1, 1)
			AssertEqual(move, false)
		})
		It("should return false as a bad actor", func() {
			g := CreateGame()
			move, _ := g.Move('q', 1, 2)
			AssertEqual(move, false)
		})
		It("should return false from low X cord", func() {
			g := CreateGame()
			move, _ := g.Move('x', 0, 1)
			AssertEqual(move, false)
		})
		It("should return false from low Y cord", func() {
			g := CreateGame()
			move, _ := g.Move('x', 1, 0)
			AssertEqual(move, false)
		})
		It("should return false high X cord", func() {
			g := CreateGame()
			move, _ := g.Move('x', 10, 1)
			AssertEqual(move, false)
		})
		It("should return false high Y cord", func() {
			g := CreateGame()
			move, _ := g.Move('x', 1, 10)
			AssertEqual(move, false)
		})
		It("should return false as same player is not allowed in a row", func() {
			g := CreateGame()
			g.Move('x', 1, 1)
			move, _ := g.Move('x', 1, 2)
			AssertEqual(move, false)
		})
	})

	Report(t)
}
