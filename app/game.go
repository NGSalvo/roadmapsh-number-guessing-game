package app

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Game struct {
	Chances     int
	Difficulty  int
	Number      int
	GuessNumber int
	Attempts    int
	Repeat      bool
}

const (
	Easy = iota + 1
	Medium
	Hard
)

func NewGame() *Game {
	return &Game{}
}

func (g *Game) getRandomNumber() int {
	return rand.IntN(100) + 1
}

func (g *Game) setAttemps() {
	switch g.Difficulty {
	case Easy:
		g.Chances = 10
	case Medium:
		g.Chances = 5
	case Hard:
		g.Chances = 3
	}
}

func (g *Game) Guess(number int) {
	g.Attempts++
	g.GuessNumber = number
}

func (g *Game) HasChances() bool {
	return g.Chances > g.Attempts
}

func (g *Game) HasWon() bool {
	return g.Number == g.GuessNumber
}

func (g *Game) reset() {
	g.Attempts = 0
}

func (g *Game) PlayAgain() bool {
	fmt.Println("Do you want to play again? (y/n)")

	var playAgain string
	fmt.Scan(&playAgain)

	for !strings.EqualFold(playAgain, "y") && !strings.EqualFold(playAgain, "n") {
		fmt.Println("Invalid input. Please enter y or n.")
		fmt.Scan(&playAgain)
	}

	if strings.EqualFold(playAgain, "n") {
		g.Repeat = false
		return false
	}

	g.reset()
	g.Repeat = true
	return true
}

func (g *Game) difficultyMenu() {

	fmt.Println("Choose a difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")
}

func (g *Game) Start() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	g.setDifficulty()

	fmt.Println("You have", g.Chances, "chances to guess the correct number.")

	g.Number = g.getRandomNumber()

	fmt.Println("Let's start the game!")

	fmt.Println("Enter your guess: ")
	var guess int
	for g.HasChances() {
		fmt.Scan(&guess)
		g.Guess(guess)
		if g.HasWon() {
			fmt.Println("Congratulations! You guessed the correct number in", g.Attempts, "attempts.")
			break
		}

		g.showHint(guess)
	}

	if !g.HasWon() {
		fmt.Println("Sorry, you ran out of chances. The correct number was", g.Number)
	}

	if !g.PlayAgain() {
		fmt.Println("Thanks for playing!")
	}
}

func (g *Game) showHint(guess int) {
	if guess > g.Number {
		fmt.Println("Incorrect! The number is less than", guess)
	} else if guess < g.Number {
		fmt.Println("Incorrect! The number is greater than", guess)
	}
}

func (g *Game) setDifficulty() {
	var difficulty int

	g.difficultyMenu()

	fmt.Scan(&difficulty)

	for difficulty < 1 || difficulty > 3 {
		fmt.Println("Invalid difficulty level. Please enter a number between 1 and 3 to select a difficulty level.")
		fmt.Println("Choose a difficulty level:")
		fmt.Scan(&difficulty)
	}
	g.Difficulty = difficulty
	g.setAttemps()
}
