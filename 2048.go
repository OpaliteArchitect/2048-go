package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var board [4][4]int

	var direction string
	turns := 0

	spawnTile(&board)
	printBoard(&board)

	for {
		fmt.Print("Enter wasd: ")
		fmt.Scan(&direction)
		fmt.Println()

		if playRound(&board, direction) {
			spawnTile(&board)
			printBoard(&board)
			turns++
		}
		fmt.Println()
	}
}

func spawnTile(board *[4][4]int) {
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

func printBoard(board *[4][4]int) {
	for i := range 4 {
		for j := range 4 {
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
}

func playRound(board *[4][4]int, direction string) bool {
	switch direction {
	case "a":
	case "w":
		rotateClockwise(board, 3)
	case "d":
		rotateClockwise(board, 2)
	case "s":
		rotateClockwise(board, 1)
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
	case "w":
		rotateClockwise(board, 1)
	case "d":
		rotateClockwise(board, 2)
	case "s":
		rotateClockwise(board, 3)
	}

	return true
}

func rotateClockwise(board *[4][4]int, times int) {
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
