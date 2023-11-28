package main

import "fmt"

func main() {
	fmt.Println("Enter a year to check the leap year: ")
	var year int
	fmt.Scanln(&year)

	leapYear(year)
}

func leapYear(calYear int) {
	if calYear%400 == 0 || (calYear%4 == 0 && calYear%100 != 0) {
		fmt.Println(true)
	} else {
		fmt.Println("Not a leap year")
	}
}
