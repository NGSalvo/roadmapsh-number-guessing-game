package main

import (
	"number-guessing-game/app"
)

func main() {

	game := app.NewGame()

	game.Start()

	for game.Repeat {
		game.Start()
	}
}
