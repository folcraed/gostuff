package main

import (
	"fmt"
	"strconv"
	"strings"
)

// toCelsius takes the supplied number and
// converts it to celsius.
func toCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) / 1.80
}

// toFahrenheit takes the supplied number and
// converts it to Fahrenheit.
func toFahrenheit(celsius float64) float64 {
	return celsius*(9/5) + 32
}

var message = "Enter \"32c\" to convert to fahrenheit, \"32f\" to convert to celsius."
var converting string

func main() {
	fmt.Println(message)
	fmt.Scanf("%s", &converting)
	var original = ""

	if strings.HasSuffix(converting, "c") {
		original = "c"
	}
	if strings.HasSuffix(converting, "f") {
		original = "f"
	}
	var newVal = strings.TrimSuffix(converting, original)
	convertMe, _ := strconv.ParseFloat(newVal, 64)

	switch original {
	case "c":
		fmt.Printf("That equals %.1f fahrenheit\n", toFahrenheit(convertMe))
	case "f":
		fmt.Printf("That equals %.1f celsius\n", toCelsius(convertMe))
	default:
		fmt.Println("Something went wrong, check your input and try again!")
	}
}
