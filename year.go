package main

import "fmt"

func year() {
	fmt.Println("Enter year: ")

	var year int
	fmt.Scanln(year)

	leapYear := yearCal(year)

	fmt.Printf("The student's letter grade is: %s\n", leapYear)

}

func yearCal(year int) bool {
	if year%4 == 0 {
		fmt.Println("This is a leap year")
	}
	return false
}
