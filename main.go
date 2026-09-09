package main

import (
	"fmt"

	"github.com/opalitearchitect/2048-go/game"
)

func main() {
	var board game.Board

	var direction string
	turns := 0

	board.SpawnTile()
	board.PrintBoard()

	for {
		fmt.Print("Enter wasd: ")
		fmt.Scan(&direction)
		fmt.Println()

		if board.PlayRound(direction) {
			board.SpawnTile()
			board.PrintBoard()
			turns++
		}
		fmt.Println()
	}
}
