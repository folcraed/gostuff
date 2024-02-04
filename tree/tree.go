// Create a ASCII tree given a height from user.
// A traditional programming exercise.
package main

import (
	"fmt"
	"strings"
)

// Set the defaults
var height = 0
var message = "\nHow high is the tree?"

func main() {
	fmt.Println(message)
	fmt.Scanf("%d", &height) // Get the starting height
	spaces := height - 1     // Set initial indentation size
	hashes := 1              // Set initial branch creation size
	stump := height - 2      // Set where the tree stump will be.

	// Create the tree using spaces and hashes
	i := 1
	for i < height {
		spacer := strings.Repeat(" ", spaces)
		hasher := strings.Repeat("#", hashes)
		fmt.Print(spacer)
		fmt.Println(hasher)
		spaces = spaces - 1
		hashes = hashes + 2
		height = height - 1
	}
	// Then give the tree it's stump
	fmt.Println(strings.Repeat(" ", stump), "#")
}
