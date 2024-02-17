// Calculates the area of circles or squares

package main

import (
	"fmt"
	"math"
)

// Circle takes the radius to calculate the area
func Circle(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

// Square takes the width & height to calculate the area
func Square(width, height float64) float64 {
	return width * height
}

var choice string
var a float64
var b float64

func main() {
	fmt.Println("What area are we calculating, (c)Circle or (s)Square?")
	fmt.Println("For circles we calculate radius, if you have diameter, enter 1/2 of the diameter.")
	fmt.Scanf("%s\n", &choice)

	// Find what to work with, circle or square
	switch choice {
	case "c":
		fmt.Print("Enter the radius: ")
		fmt.Scanf("%f", &a)
	case "s":
		fmt.Print("Enter the width and height (separate with space): ")
		fmt.Scanf("%f %f", &a, &b)
	default:
		fmt.Println("Something went wrong, check you input and try again!")
	}

	// Depending on whether circle or square, do the calculation
	// and print the result.
	if choice == "c" {
		fmt.Printf("The area of the circle is %.2f\n", Circle(a))
	} else if choice == "s" {
		fmt.Printf("The area of the square is %.2f\n", Square(a, b))
	} else {
		fmt.Println("Something's wrong, can not calculate!")
	}
}
