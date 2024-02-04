// Take a user supplied list of numbers and average them.
package main

import "fmt"

var message = "How many numbers do you need to average?"

func main() {
	// Set default values
	size := 0
	added := 0.0

	// Print message and get size of slice from user
	fmt.Println(message)
	fmt.Scanln(&size)

	// Create the slice and get the numbers from user
	input := make([]float64, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%f", &input[i])
	}
	// Add the numbers together
	for i := 0; i < size; i++ {
		added += input[i]
	}
	// Calculate and print the results.
	divisor := float64(size)
	fmt.Printf("The average of those numbers is %.3f\n", (added / divisor))
}
