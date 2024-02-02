package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toKilometers(miles float64) float64 {
	return miles * 1.60934
}

func toMiles(kilometers float64) float64 {
	return kilometers / 1.60934
}

var message = "Enter \"32m\" to convert miles, \"32k\" to convert kilometers."
var converting string

func main() {
	fmt.Println(message)
	fmt.Scanf("%s", &converting)
	var original = ""

	if strings.HasSuffix(converting, "m") {
		original = "m"
	}
	if strings.HasSuffix(converting, "k") {
		original = "k"
	}

	var newVal = strings.TrimSuffix(converting, original)
	convertMe, _ := strconv.ParseFloat(newVal, 64)

	switch original {
	case "m":
		fmt.Printf("That equals %.2f kilometers.\n", toKilometers(convertMe))
	case "k":
		fmt.Printf("That equals %.2f miles.\n", toMiles(convertMe))
	default:
		fmt.Println("Something went wrong, check your input and try again!")
	}
}
