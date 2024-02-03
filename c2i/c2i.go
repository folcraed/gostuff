package main

import (
	"fmt"
	"strconv"
	"strings"
)

// toCentimeters takes a value passed to it and
// converts it to centimeters.
func toCentimeters(inches float64) float64 {
	return inches / 0.39370
}

// toInches takes a value passed to it and
// converts it to inches.
func toInches(centimeters float64) float64 {
	return centimeters * 0.39370
}

var message = "Enter a value as \"32c\" to convert from cm, or \"32i\" to convert from inches.\n"
var text string

func main() {
	fmt.Println(message)
	fmt.Scanf("%s", &text)
	var original = ""

	if strings.HasSuffix(text, "c") {
		original = "c"
	}
	if strings.HasSuffix(text, "i") {
		original = "i"
	}

	var newVal = strings.TrimSuffix(text, original)
	converting, _ := strconv.ParseFloat(newVal, 64)

	switch original {
	case "c":
		fmt.Printf("That equals %.2f inches\n", toInches(converting))
	case "i":
		fmt.Printf("That equals %.2f centimeters\n", toCentimeters(converting))
	default:
		fmt.Println("Something's wrong, check your input and try again!")
	}
}
