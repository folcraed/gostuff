// A file for experimenting with things

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

var empties = []string{" .", " .", " .", " .", " .", " .", " .", " .", " .", " .", " .", " ."}
var ships = []string{" ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}
var guess string
var count = 0
var SHIPS = 0

const SUNK = " X"

const MISSED = " *"

func sank() {
	fmt.Println("You sank my battleship!")
	SHIPS += 1
}

func missed() {
	fmt.Println("Missed!")
	count += 1
}

func displayBoard() {
	fmt.Println("   A  B  C  D")
	fmt.Printf("1 %s %s %s %s\n", empties[0], empties[3], empties[6], empties[9])
	fmt.Printf("2 %s %s %s %s\n", empties[1], empties[4], empties[7], empties[10])
	fmt.Printf("3 %s %s %s %s\n", empties[2], empties[5], empties[8], empties[11])
}

func getRandom() {
	rand.NewSource(time.Now().UnixNano())
	p := rand.Perm(12)
	for _, r := range p[:3] {
		ships[r] = SUNK
	}
}

func getGuess() {
	fmt.Print("Enter a guess: ")
	fmt.Scanln(&guess)
}

func main() {
	getRandom()
	fmt.Println()
	fmt.Println("This is the game of Battleships.")
	fmt.Println("Enter coordinates to shoot (like 'a1' or 'c3').")
	fmt.Println("Enter 'x' to quit early")
	fmt.Println()
	for {
		if SHIPS == 3 {
			fmt.Println("You sunk ALL my battleships!")
			displayBoard()
			fmt.Printf("It took you %d tries.\n", count)
			break
		}
		displayBoard()
		getGuess()
		switch guess {
		case "a1":
			if ships[0] == SUNK {
				empties[0] = SUNK
				sank()
			} else {
				missed()
				empties[0] = MISSED
			}
		case "a2":
			if ships[1] == SUNK {
				empties[1] = SUNK
				sank()
			} else {
				missed()
				empties[1] = MISSED
			}
		case "a3":
			if ships[2] == SUNK {
				empties[2] = SUNK
				sank()
			} else {
				missed()
				empties[2] = MISSED
			}
		case "b1":
			if ships[3] == SUNK {
				empties[3] = SUNK
				sank()
			} else {
				missed()
				empties[3] = MISSED
			}
		case "b2":
			if ships[4] == SUNK {
				empties[4] = SUNK
				sank()
			} else {
				missed()
				empties[4] = MISSED
			}
		case "b3":
			if ships[5] == SUNK {
				empties[5] = SUNK
				sank()
			} else {
				missed()
				empties[5] = MISSED
			}
		case "c1":
			if ships[6] == SUNK {
				empties[6] = SUNK
				sank()
			} else {
				missed()
				empties[6] = MISSED
			}
		case "c2":
			if ships[7] == SUNK {
				empties[7] = SUNK
				sank()
			} else {
				missed()
				empties[7] = MISSED
			}
		case "c3":
			if ships[8] == SUNK {
				empties[8] = SUNK
				sank()
			} else {
				missed()
				empties[8] = MISSED
			}
		case "d1":
			if ships[9] == SUNK {
				empties[9] = SUNK
				sank()
			} else {
				missed()
				empties[9] = MISSED
			}
		case "d2":
			if ships[10] == SUNK {
				empties[10] = SUNK
				sank()
			} else {
				missed()
				empties[10] = MISSED
			}
		case "d3":
			if ships[11] == SUNK {
				empties[11] = SUNK
				sank()
			} else {
				missed()
				empties[11] = MISSED
			}
		case "x":
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}
}
