package main

type game struct {
}

type Board [][]bool

func (game *game) initialize(board Board) {
}

func (game *game) advanceOneTick() {

}

func (gam *game) getCurrentBoard() [3][3]bool {
	return [3][3]bool{
		{false, false, false},
		{true, true, true},
		{false, false, false},
	}
}

func Game() *game {
	return &game{}
}
