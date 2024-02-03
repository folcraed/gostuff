// Take a user supplied list of numbers and average them.
package main

import "fmt"

var message = "How many numbers do you need to average?"

func main() {
	size := 0
	added := 0.0
	fmt.Println(message)
	fmt.Scanln(&size)
	input := make([]float64, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter a number: ")
		fmt.Scanf("%f", &input[i])
	}
	for i := 0; i < size; i++ {
		added += input[i]
	}
	// fmt.Println(input)
	divisor := float64(size)
	fmt.Printf("The average of those numbers is %.3f\n", (added / divisor))
}
