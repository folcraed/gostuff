package main

import (
	"fmt"
	"strconv"
	"strings"
)

// toPounds takes input and converts ot to pounds.
func toPounds(weight float64) float64 {
	return weight * 2.2046
}

// toKilograms takes input and converts it to kilograms.
func toKilograms(weight float64) float64 {
	return weight / 2.2046
}

var message = "Enter weight as \"32p\" to convert pounds, or \"32k\" to convert kilograms"
var converting string

func main() {
	fmt.Println(message)
	fmt.Scanf("%s", &converting)
	var original = ""

	if strings.HasSuffix(converting, "p") {
		original = "p"
	}
	if strings.HasSuffix(converting, "k") {
		original = "k"
	}

	var newVal = strings.TrimSuffix(converting, original)
	convertMe, _ := strconv.ParseFloat(newVal, 64)

	switch original {
	case "p":
		fmt.Printf("That equals %.2f kilograms.\n", toKilograms(convertMe))
	case "k":
		fmt.Printf("That equals %.2f pounds.\n", toPounds(convertMe))
	default:
		fmt.Println("Something went wrong, check you input and try again!")
	}
}
