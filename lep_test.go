package main

import (
	"testing"
)

func TestInitialGame(t *testing.T) {
	game := Game()
	if game == nil {
		t.Errorf("Wah wah no game")
	}
}

func TestCanInitializeGame(t *testing.T) {
	game := Game()

	board := [2][2]bool{
		{true,false},
		{false,true},
	}

	game.initialize(board)
}