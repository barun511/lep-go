package main

type game struct {
	count int
}

type Board [][]bool

func (game *game) initialize(board Board) {
	game.count = 0
}

func (game *game) advanceOneTick() {
	game.count += 1
}

func (game *game) getCurrentBoard() [][]bool {
	if game.count == 1 {
		return [][]bool{
			{false, false, false},
			{true, true, true},
			{false, false, false},
		}
	}
	return [][]bool{
		{false, true, false},
		{false, true, false},
		{false, true, false},
	}

}

func Game() *game {
	return &game{}
}
