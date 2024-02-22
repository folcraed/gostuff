// A file for experimenting with things

package main

import (
	"fmt"
)

// TODO: Would like to move this employee info into another file,
// something read/write that holds the data.
type Employee struct {
	Name  string
	Title string
}

var employees = [6]Employee{
	{"Bob", "boss"},
	{"Jeremy", "installer"},
	{"Kevin", "towman"},
	{"BJ", "towman"},
	{"Dusty", "towman"},
	{"Jay", "towman"},
}

var searchFor = ""
var listing = ""

func lookUp() {
	fmt.Print("Enter the name you're searching for: ")
	fmt.Scanln(&searchFor)
}

func showEmpl() {
	for i := 0; i < len(employees); i++ {
		fmt.Println(employees[i].Name, "is", employees[i].Title)
	}

}

func showInfo() {
	for j := 0; j < len(employees); j++ {
		if searchFor == "" {
			break
		}
		if employees[j].Name == searchFor {
			switch employees[j].Title {
			case "boss":
				fmt.Println(employees[j].Name, "is", employees[j].Title)
			case "installer":
				fmt.Println(employees[j].Name, "is", employees[j].Title)
			case "towman":
				fmt.Println(employees[j].Name, "is", employees[j].Title)
			}
			break
		}
		if employees[j].Name != searchFor && j == len(employees)-1 {
			fmt.Println("Sorry", searchFor, "doesn't work here.")
		}
	}
}

func main() {
	// TODO: Would like to make this into a menu system, presenting the user
	// with choices on what they want to do.

	lookUp()
	showInfo()
	fmt.Println("Would you like to see a list of employees? (y/n)")
	fmt.Scanln(&listing)
	if listing == "y" {
		showEmpl()
	}
	fmt.Println()
	fmt.Println("Bye for now!")
}
