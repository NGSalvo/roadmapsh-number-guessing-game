package tests

import (
	"bytes"
	"io"
	"number-guessing-game/app"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("game should be created with chosen difficulty level", func(t *testing.T) {

		tests := []struct {
			name       string
			difficulty string
			expected   *app.Game
		}{
			{"game should be created with difficulty level easy", "easy", &app.Game{Chances: 10, Difficulty: "easy"}},
			{"game should be created with default difficulty level medium", "medium", &app.Game{Chances: 5, Difficulty: "medium"}},
			{"game should be created with default difficulty level hard", "hard", &app.Game{Chances: 3, Difficulty: "hard"}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				game := app.NewGame(tt.difficulty)
				asserts.Equal(tt.expected.Difficulty, game.Difficulty)
				asserts.Equal(tt.expected.Chances, game.Chances)
			})
		}
	})

	t.Run("game should be won when guessing the right number withing the number of chances", func(t *testing.T) {

		game := app.NewGame("easy")
		game.Number = 50
		game.Guess(50)
		asserts.Equal(1, game.Attempts)
		asserts.Equal(50, game.GuessNumber)
		asserts.True(game.HasWon())
	})

	t.Run("game should be not won when guessing the wrong number withing the number of chances", func(t *testing.T) {

		game := app.NewGame("easy")
		game.Number = 50
		for game.HasChances() {
			game.Guess(25)
		}
		asserts.Equal(10, game.Attempts)
		asserts.False(game.HasWon())
	})

	t.Run("when game starts, it should display a welcome message along with the rules of the game", func(t *testing.T) {
		game := app.NewGame("easy")

		output := outputToString(game.Start)
		asserts.Equal("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 10 chances to guess the correct number.\n", output)
	})
}

func outputToString(callback func()) string {

	// Create a pipe to capture the output
	r, w, _ := os.Pipe()

	// Save the original stdout
	oldStdout := os.Stdout

	// Assign the write end of the pipe to stdout
	os.Stdout = w

	// Ensure that stdout is restored after the test
	defer func() {
		os.Stdout = oldStdout
		w.Close()
	}()

	// Call the function or code that prints to stdout
	callback()

	// Close the write end of the pipe to signal EOF
	w.Close()

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
