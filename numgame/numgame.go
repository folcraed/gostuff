// Traditional number guessing game exercise.
package main

import (
	"fmt"
	"math/rand"
)

// Setup the starting defaults.
var secretNum = rand.Intn(100)
var guesses = 10
var count = 0
var guess = 0

func main() {
	fmt.Println("Welcome to the number guessing game!\nIt randomly picks a number between 1-100 for you to guess.\nYou have 10 tries to guess the number.\n")
	for i := 1; i <= guesses; i++ {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)
		count = count + 1
		left := guesses - count
		if i == 10 {
			fmt.Printf("Sorry, you ran out of guesses. The number was %d\n", secretNum)
			break
		} else if guess < secretNum {
			fmt.Printf("You guessed to low, you have %d guesses left. Guess again.\n", left)
		} else if guess > secretNum {
			fmt.Printf("You guessed to high, you have %d guesses left. Guess again.\n", left)
		} else {
			fmt.Printf("Correct! The number was %d\n", secretNum)
			break
		}

	}
}
