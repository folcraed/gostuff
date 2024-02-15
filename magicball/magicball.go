// The "Magic 8_ball" game

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
)

// The answers the 8-ball might give
var answers = [21]string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes, definitely",
	"You may rely on it",
	"As I see it, yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy, try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
	"You seriously want me to answer?",
}

// This function gets the question from the user,
// checking if they entered a string. If they did,
// it returns control to main. It doesn't do anything
// with the string, the string is actually not needed.
// If they just pressed ENTER, it exits the program.
func getQuestion() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if line == "\n" {
		fmt.Println("The Magic 8-Ball has left!")
		os.Exit(0)
	}
}

func main() {
	for i := 0; i < 1000; i++ { // Create a loop that should never be met.
		numAnswer := rand.Intn(21) // Randomly pick an answer.
		fmt.Println("Ask the Magic 8-Ball a question, <ENTER> to exit.")
		getQuestion() // Note the function doesn't take any parameters, nor returns any.
		theAnswer := answers[numAnswer]
		fmt.Println(theAnswer)
	}
}
