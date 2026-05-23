package main

type game struct {
	count            int
	isUnderpopulated bool
	staySame         bool
	board            Board
}

type Board [][]bool

func (game *game) initialize(board Board) {
	game.board = board
}

type Delta struct {
	row int
	col int
}

func print_board(board Board) {
	for i := range board {
		for j := range board[i] {
			print(board[i][j])
			print(" ")
		}
		println()
	}
	println()
}

func (game *game) advanceOneTick() {
	newBoard := make([][]bool, len(game.board))
	for i := range game.board {
		newBoard[i] = make([]bool, len(game.board[i]))
		copy(newBoard[i], game.board[i])
	}

	deltas := [][]int{{-1, 0},
		{-1, -1}, {-1, 1},
		{0, -1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

	for i := range game.board {
		for j := range game.board[i] {
			count := 0
			for delta_i := range deltas {
				delta := deltas[delta_i]
				if i+delta[0] < 0 || i+delta[0] >= len(game.board) ||
					j+delta[1] < 0 || j+delta[1] >= len(game.board[0]) {
					continue
				}
				if game.board[i+delta[0]][j+delta[1]] {
					count += 1
				}
			}

			if game.board[i][j] == false {
				if count == 3 {
					newBoard[i][j] = true
				}
			} else {
				if count == 2 || count == 3 {
					newBoard[i][j] = true
				} else {
					newBoard[i][j] = false
				}
			}

		}
	}

	// print_board(newBoard)

	game.board = newBoard
}

func (game *game) getCurrentBoard() [][]bool {
	return game.board

}

func Game() *game {
	return &game{}
}
