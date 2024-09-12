package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	game := NewGame("easy")

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have", game.Chances, "chances to guess the correct number.")

	game.Number = 5

	fmt.Println("Let's start the game!")
	var guess int
	for game.HasChances() {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		game.Guess(guess)
		if game.IsWon() {
			fmt.Println("Congratulations! You guessed the correct number in", game.Attempts, "attempts.")
			break
		}
	}

	if !game.IsWon() {
		fmt.Println("Sorry, you ran out of chances. The correct number was", game.Number)
	}

	fmt.Println("Thanks for playing!")
}

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

func (g *Game) IsWon() bool {
	return g.Number == g.GuessNumber
}
