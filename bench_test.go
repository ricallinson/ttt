package ttt

import (
	"testing"
)

func BenchmarkMove(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g := CreateGame()
		g.Move('x', 1, 1)
		g.Move('o', 1, 3)
		g.Move('x', 2, 1)
		g.Move('o', 2, 3)
		g.Move('x', 3, 1)
	}
}
