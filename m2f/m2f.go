package main

import (
	"fmt"
	"strconv"
	"strings"
)

func toFeet(distance float64) float64 {
	return distance * 3.2808
}

func toMeters(distance float64) float64 {
	return distance * 0.3048
}

var message = "Enter a value as \"12m\" to convert from meters, \"12f\" to convert from feet."
var from string

func main() {
	fmt.Println(message)
	fmt.Scanf("%s", &from)
	var original = ""

	if strings.HasSuffix(from, "m") {
		original = "m"
	}
	if strings.HasSuffix(from, "f") {
		original = "f"
	}

	var newVal = strings.TrimSuffix(from, original)
	converting, _ := strconv.ParseFloat(newVal, 64)

	switch original {
	case "m":
		fmt.Printf("That equals %.2f feet\n", toFeet(converting))
	case "f":
		fmt.Printf("That equals %.2f meters\n", toMeters(converting))
	default:
		fmt.Println("Something's wrong, check your input and try again!")
	}
}
