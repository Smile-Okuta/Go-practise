package main

import "fmt"

func main() {
	fmt.Println("Enter your age: ")
	var age int

	fmt.Scanln(&age)

	price := ageTicket(age)

	fmt.Printf("Your ticket price is: %v\n", price)

}

func ageTicket(num int) string {

	switch {
	case num <= 12:
		return "N5"
	case num > 12 && num <= 64:
		return "N10"
	case num >= 65:
		return "N7"
	default:
		return "Enter a valid age"
	}

}
