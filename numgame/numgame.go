// Traditional number guessing game exercise.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Setup the starting defaults.
var guesses int = 10
var count int = 0
var guess int = 0

func main() {
	rand.NewSource(time.Now().UnixNano())
	secretNum := rand.Intn(100)

	fmt.Println("Welcome to the number guessing game!\nIt randomly picks a number between 1-100 for you to guess.\nYou have 10 tries to guess the number.")
	fmt.Println()
	for i := 1; i <= guesses; i++ {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid guess: err:", err)
			continue
		}
		count += 1
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
