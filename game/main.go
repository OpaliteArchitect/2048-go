package game

type game struct {
	board Board
	turns int
}

type Game interface {
	Move(direction Direction) bool
	PrintBoard()
	GetTurns() int
	GetBoard() Board
}

func (game *game) Move(direction Direction) bool {
	moved := game.board.playRound(direction)
	if moved {
		game.turns++
		game.board.SpawnTile()
	}

	return moved
}

func New() game {
	game := game{
		board: Board{},
		turns: 0,
	}

	game.board.SpawnTile()
	game.board.SpawnTile()

	return game
}

func (game *game) PrintBoard() {
	game.board.PrintBoard()
}

func (game *game) GetTurns() int {
	return game.turns
}

func (game *game) GetBoard() Board {
	return game.board
}
