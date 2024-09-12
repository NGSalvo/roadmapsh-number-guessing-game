package app

import (
	"fmt"
	"math/rand/v2"
)

type Game struct {
	Chances     int
	Difficulty  string
	Number      int
	GuessNumber int
	Attempts    int
}

func NewGame(difficulty string) *Game {
	game := Game{
		Difficulty: difficulty,
	}
	game.setAttemps()
	game.Number = game.getRandomNumber()
	return &game
}

func (g *Game) getRandomNumber() int {
	return rand.IntN(100) + 1
}

func (g *Game) setAttemps() {
	switch g.Difficulty {
	case "easy":
		g.Chances = 10
	case "medium":
		g.Chances = 5
	case "hard":
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

func (g *Game) Start() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have", g.Chances, "chances to guess the correct number.")

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
	}

	if !g.HasWon() {
		fmt.Println("Sorry, you ran out of chances. The correct number was", g.Number)
	}
	fmt.Println("Thanks for playing!")
}
