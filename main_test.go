package main

import (
	"number-guessing-game/app"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("app.Game should be created with chosen difficulty level", func(t *testing.T) {

		tests := []struct {
			name       string
			difficulty string
			expected   *app.Game
		}{
			{"app.Game should be created with difficulty level easy", "easy", &app.Game{Chances: 10, Difficulty: "easy"}},
			{"app.Game should be created with default difficulty level medium", "medium", &app.Game{Chances: 5, Difficulty: "medium"}},
			{"app.Game should be created with default difficulty level hard", "hard", &app.Game{Chances: 3, Difficulty: "hard"}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				game := app.NewGame(tt.difficulty)
				asserts.Equal(tt.expected.Difficulty, game.Difficulty)
				asserts.Equal(tt.expected.Chances, game.Chances)
			})
		}
	})

	t.Run("app.Game should won when guessing the right number withing the number of chances", func(t *testing.T) {

		game := app.NewGame("easy")
		game.Number = 50
		game.Guess(50)
		asserts.Equal(1, game.Attempts)
		asserts.Equal(50, game.GuessNumber)
		asserts.True(game.HasWon())
	})

	t.Run("app.Game should not won when guessing the wrong number withing the number of chances", func(t *testing.T) {

		game := app.NewGame("easy")
		game.Number = 50
		for game.HasChances() {
			game.Guess(25)
		}
		asserts.Equal(10, game.Attempts)
		asserts.False(game.HasWon())
	})
}
