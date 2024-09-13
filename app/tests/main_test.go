package tests

import (
	"bytes"
	"io"
	"number-guessing-game/app"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	asserts := assert.New(t)

	t.Run("game should be created with chosen difficulty level", func(t *testing.T) {

		tests := []struct {
			name       string
			difficulty int
			expected   *app.Game
		}{
			{"game should be created with difficulty level easy", 1, &app.Game{Chances: 10, Difficulty: 1}},
			{"game should be created with default difficulty level medium", 2, &app.Game{Chances: 5, Difficulty: 2}},
			{"game should be created with default difficulty level hard", 3, &app.Game{Chances: 3, Difficulty: 3}},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				game := app.NewGame()
				asserts.Equal(tt.expected.Difficulty, game.Difficulty)
				asserts.Equal(tt.expected.Chances, game.Chances)
			})
		}
	})

	t.Run("game should be won when guessing the right number withing the number of chances", func(t *testing.T) {

		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10
		game.Number = 50
		game.Guess(50)
		asserts.Equal(1, game.Attempts)
		asserts.Equal(50, game.GuessNumber)
		asserts.True(game.HasWon())
	})

	t.Run("game should be not won when guessing the wrong number withing the number of chances", func(t *testing.T) {

		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10
		game.Number = 50
		for game.HasChances() {
			game.Guess(25)
		}
		asserts.Equal(10, game.Attempts)
		asserts.False(game.HasWon())
	})

	t.Run("when game starts, it should display a welcome message along with the rules of the game", func(t *testing.T) {
		t.Skip("need to fix")
		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10

		output := outputToString(game.Start)
		asserts.Contains(output, "Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 10 chances to guess the correct number.\n")
	})

	t.Run("should be able to input a guess number", func(t *testing.T) {
		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10

		// Backup the original stdin
		oldStdin := os.Stdin

		// Restore stdin after test
		defer func() { os.Stdin = oldStdin }()

		// Create a pipe to simulate stdin
		r, w, _ := os.Pipe()

		// Set the pipe reader as stdin
		os.Stdin = r

		// Write input to the pipe writer
		input := "1 50"
		w.WriteString(input)
		w.Close() // Close the writer to simulate end of input

		expectedInput, _ := strconv.Atoi(input)

		// Call the function or code that reads from stdin
		output := outputToString(game.Start)

		asserts.Contains(output, "Enter your guess: ")
		asserts.Equal(expectedInput, game.GuessNumber)
	})

	t.Run("should greet the user when the game is won", func(t *testing.T) {
		t.Skip("need to fix because theres a random number generated")
		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10
		game.Number = 50

		// Backup the original stdin
		oldStdin := os.Stdin

		// Restore stdin after test
		defer func() { os.Stdin = oldStdin }()

		// Create a pipe to simulate stdin
		r, w, _ := os.Pipe()

		// Set the pipe reader as stdin
		os.Stdin = r

		// Write input to the pipe writer
		input := "1 50"
		w.WriteString(input)
		w.Close() // Close the writer to simulate end of input

		output := outputToString(game.Start)

		asserts.True(game.HasWon())
		asserts.Contains(output, "Congratulations! You guessed the correct number in 1 attempts.")
		asserts.Equal(50, game.GuessNumber)
	})

	t.Run("should greet the user when the game is lost", func(t *testing.T) {
		t.Skip("need to fix because theres a random number generated")
		game := app.NewGame()
		game.Difficulty = 1
		game.Chances = 10
		game.Number = 50
		for game.HasChances() {
			game.Guess(25)
		}
		asserts.Equal(10, game.Attempts)
		asserts.False(game.HasWon())
		asserts.Contains(outputToString(game.Start), "Sorry, you ran out of chances. The correct number was 50")
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
