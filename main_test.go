package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("Game should be created with default difficulty level", func(t *testing.T) {

		tests := []struct {
			name       string
			difficulty string
			expected   *Game
		}{
			{"Game should be created with difficulty level easy", "easy", &Game{Chances: 10, Difficulty: "easy"}},
			{"Game should be created with default difficulty level medium", "medium", &Game{Chances: 5, Difficulty: "medium"}},
			{"Game should be created with default difficulty level hard", "hard", &Game{Chances: 3, Difficulty: "hard"}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				game := NewGame(tt.difficulty)
				asserts.Equal(tt.expected.Difficulty, game.Difficulty)
				asserts.Equal(tt.expected.Chances, game.Chances)
			})
		}
	})

	t.Run("Game should won when guessing the right number withing the number of chances", func(t *testing.T) {

		game := NewGame("easy")
		game.Number = 50
		game.Guess(50)
		asserts.Equal(1, game.Attempts)
		asserts.Equal(50, game.GuessNumber)
		asserts.True(game.IsWon())
	})

	t.Run("Game should not won when guessing the wrong number withing the number of chances", func(t *testing.T) {

		game := NewGame("easy")
		game.Number = 50
		for game.HasChances() {
			game.Guess(25)
		}
		asserts.Equal(10, game.Attempts)
		asserts.False(game.IsWon())
	})
}
