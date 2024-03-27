/*
   Takes a user supplied starting date and asks for how many years, months and days
   into the past or future to calculate from the starting date. Returns the year,
   month and day of the difference.
*/

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// User prompts with instructions
var message = "Enter the first date as '2023-06-12' (Year-Month-Day).\nNote single-digit months and days MUST have '0' prefix."
var getDiff = `Now enter the time difference, '23 8 6' to calculate
into the future, '-22 -5 -10' to calculate into the past.
You can also enter full dates like '1959 10 5' to get how
many years, months and days are between them.`

func main() {
	var timeInput string
	fmt.Println(message)
	fmt.Println()
	fmt.Scanf("%s", &timeInput) // Get the starting date from the user
	theFirst, err := time.Parse(time.DateOnly, timeInput)
	if err != nil {
		fmt.Println("You entered the wrong date format:", err)
		os.Exit(1) // If they enter the wrong format, exit the program.
	}

	// Variables to hold the difference input
	var year int
	var month int
	var day int

	// Get the date difference and store it
	fmt.Println()
	fmt.Println(getDiff)
	fmt.Println()
	fmt.Scan(&year, &month, &day)

	// Take the difference and add or subtract it from the original date
	difference := theFirst.AddDate(year, month, day)

	// Convert the time format to a string
	toString := difference.String()

	// Strip off the time format parts we don't want
	final := strings.TrimSuffix(toString, "00:00:00 +0000 UTC")

	// Give the user the answer
	fmt.Println("\nThe difference is", final)
}
