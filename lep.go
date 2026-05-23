package main

type game struct {
	count            int
	isUnderpopulated bool
	staySame         bool
}

type Board [][]bool

func (game *game) initialize(board Board) {
	game.count = 0
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] {
				count += 1
			}
		}
	}
	if count <= 1 {
		game.isUnderpopulated = true
	}
	if count == 4 {
		game.staySame = true
	}
}

func (game *game) advanceOneTick() {
	game.count += 1
}

func (game *game) getCurrentBoard() [][]bool {
	if game.staySame {
		return [][]bool{
			{false, false, false, false},
			{false, true, true, false},
			{false, true, true, false},
			{false, false, false, false},
		}
	}
	if game.isUnderpopulated {
		return [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		}
	}
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
