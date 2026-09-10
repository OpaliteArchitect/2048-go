package main

import (
	"fmt"

	"github.com/opalitearchitect/2048-go/game"
)

func main() {
	game := game.New()

	for {
		fmt.Println("Turns:", game.GetTurns())
		game.PrintBoard()

		fmt.Print("Enter wasd: ")
		var input string
		fmt.Scan(&input)
		fmt.Println()

		direction, ok := parseDirection(input)
		if !ok {
			continue
		}

		game.Move(direction)
	}
}

func parseDirection(input string) (game.Direction, bool) {
	switch input {
	case "w", "W":
		return game.Up, true
	case "a", "A":
		return game.Left, true
	case "s", "S":
		return game.Down, true
	case "d", "D":
		return game.Right, true
	default:
		return 0, false
	}
}
