package main

import (
	"math/rand/v2"
	"reflect"
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

	board := make([][]bool, rand.IntN(5)+2)

	for i := range board {
		board[i] = make([]bool, rand.IntN(5)+2)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = rand.IntN(1)%2 == 0
		}
	}

	game.initialize(board)
}

func TestCanInitializeGameWithSpecificBoardSize(t *testing.T) {
	game := Game()

	board := [][]bool{
		{false, true, false},
		{false, true, false},
		{false, true, false},
	}

	game.initialize(board)
}

func TestCanAdvanceBlinkerOneTick(t *testing.T) {
	game := Game()

	board := [][]bool{
		{false, true, false},
		{false, true, false},
		{false, true, false},
	}

	game.initialize(board)

	game.advanceOneTick()

	newBoard := game.getCurrentBoard()

	if !reflect.DeepEqual(newBoard, [][]bool{
		{false, false, false},
		{true, true, true},
		{false, false, false},
	}) {
		t.Errorf("Board does not tick correctly")
	}

	game.advanceOneTick()

	newBoard = game.getCurrentBoard()

	if !reflect.DeepEqual(newBoard, [][]bool{
		{false, true, false},
		{false, true, false},
		{false, true, false},
	}) {
		t.Errorf("Board still does not tick correctly")
	}
}



func TestUnderpopulationSingleCell(t *testing.T) {
	game := Game()

	board := [][]bool{
		{false, false, false},
		{false, true, false},
		{false, false, false},
	}

	game.initialize(board)

	game.advanceOneTick()

	newBoard := game.getCurrentBoard()

	if !reflect.DeepEqual(newBoard, [][]bool{
		{false, false, false},
		{false, false, false},
		{false, false, false},
	}) {
		t.Errorf("Board does not tick correctly")
	}

	for i := 0; i<= 10; i++ {
		game.advanceOneTick()
		newBoard = game.getCurrentBoard()

		if !reflect.DeepEqual(newBoard, [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}) {
			t.Errorf("Board still does not tick correctly")
		}
	}
}



func TestStillLife(t *testing.T) {
	game := Game()

	board := [][]bool{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}

	game.initialize(board)

	for i := 0; i<= 10; i++ {
		game.advanceOneTick()
		newBoard := game.getCurrentBoard()

		if !reflect.DeepEqual(newBoard, [][]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, true, true, false},
			{false, false, false, false},
		}) {
			t.Errorf("Board still does not tick correctly")
		}
	}
}

func TestFourDeath(t *testing.T) {
	game := Game()

	board := [][]bool{
		{true, false, false, true},
		{false, false, false, false},
		{false, false, false, false},
		{true, false, false, true},
	}

	game.initialize(board)

	for i := 0; i<= 10; i++ {
		game.advanceOneTick()
		newBoard := game.getCurrentBoard()

		if !reflect.DeepEqual(newBoard, [][]bool{
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
			{false, false, false, false},
		}) {
			t.Errorf("Board still does not tick correctly")
		}
	}
}


