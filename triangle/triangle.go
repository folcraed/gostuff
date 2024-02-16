// Calculates the length of distance between two end points of
// a right-angle triangle. Returns the distance with .2 accuracy.

package main

import (
	"fmt"
	"math"
)

func main() {
	var side1 float64
	var side2 float64

	fmt.Print("Enter the first side length: ")
	fmt.Scanln(&side1)
	fmt.Print("Enter the second side length: ")
	fmt.Scanln(&side2)

	side3 := (side1 * side1) + (side2 * side2)
	finalValue := math.Sqrt(side3)
	fmt.Printf("The length of the 3rd side is %.2f\n", finalValue)
}
