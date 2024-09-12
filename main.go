package main

import (
	"fmt"
	"number-guessing-game/app"
)

func main() {

	game := app.NewGame("easy")

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
		if game.HasWon() {
			fmt.Println("Congratulations! You guessed the correct number in", game.Attempts, "attempts.")
			break
		}
	}

	if !game.HasWon() {
		fmt.Println("Sorry, you ran out of chances. The correct number was", game.Number)
	}

	fmt.Println("Thanks for playing!")
}
