package main

import (
	"math/rand/v2"
	"testing"
)

func TestInitialGame(t *testing.T) {
	game := Game()
	if game == nil {
		t.Errorf("Wah wah no game")
	}
}

func TestCanInitializeGameWithRandomBoardSize(t *testing.T) {
	game := Game()

	board := make([][]bool, rand.IntN(5) + 2)

	for i := range board {
		board[i] = make([]bool, rand.IntN(5) + 2)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = rand.IntN(1) % 2 == 0;
		}
	}

	game.initialize(board)
}

