package main

import "fmt"

func main() {

	fmt.Println("Enter your alphabet Character: ")
	var alpha string
	fmt.Scanln(&alpha)

	alphabet(alpha)

}

func alphabet(vow string) string {

	switch vow {
	case "A", "E", "I", "O", "U":
		fmt.Println("This is a vowel")
	default:
		fmt.Println("This is a Consonant")
	}
	return vow
}
