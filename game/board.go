package game

import (
	"fmt"
	"math/rand/v2"
)

type board [4][4]int

type Direction int

const (
	Up Direction = iota
	Left
	Right
	Down
)

func (board *board) SpawnTile() {
	for {
		i, j := rand.IntN(4), rand.IntN(4)
		if board[i][j] == 0 {
			if rand.IntN(100) < 90 {
				board[i][j] = 2
			} else {
				board[i][j] = 4
			}
			break
		}
	}
}

func (board *board) PrintBoard() {
	for i := range 4 {
		for j := range 4 {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func (board *board) playRound(direction Direction) bool {
	switch direction {
	case Left:
	case Up:
		board.rotateClockwise(3)
	case Right:
		board.rotateClockwise(2)
	case Down:
		board.rotateClockwise(1)
	default:
		return false
	}

	for i := range 4 {
		for j := 1; j < 4; j++ {
			for j > 0 && board[i][j] != 0 && board[i][j-1] == 0 {
				board[i][j-1] = board[i][j]
				board[i][j] = 0
				j--
			}
		}
	}

	for i := range 4 {
		for j := range 3 {
			if board[i][j] == board[i][j+1] {
				board[i][j] *= 2
				board[i][j+1] = 0
			}
		}
	}

	for i := range 4 {
		for j := 1; j < 4; j++ {
			for j > 0 && board[i][j] != 0 && board[i][j-1] == 0 {
				board[i][j-1] = board[i][j]
				board[i][j] = 0
				j--
			}
		}
	}

	switch direction {
	case Up:
		board.rotateClockwise(1)
	case Right:
		board.rotateClockwise(2)
	case Down:
		board.rotateClockwise(3)
	}

	return true
}

func (board *board) rotateClockwise(times int) {
	for range times {
		for i := range 4 {
			for j := i + 1; j < 4; j++ {
				board[i][j], board[j][i] = board[j][i], board[i][j]
			}
		}

		for i := range 4 {
			for j := range 2 {
				board[i][j], board[i][3-j] = board[i][3-j], board[i][j]
			}
		}
	}
}
