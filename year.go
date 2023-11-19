package main

import "fmt"

func year() {
	fmt.Println("Enter year: ")

	var year int
	fmt.Scanln(year)

	yearCal(year)

}

func yearCal(year int) bool {
	if year%4 == 0 {
		fmt.Println("This is a leap year")
	}
	return false
}
