package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter year: ")

	var year int
	fmt.Scanln(&year)

	yearCal(year)

}

func yearCal(year int) string {
	if year%4 == 0 {
		fmt.Println("This is a leap year")

	}

	return "This is not a leap year"
}
