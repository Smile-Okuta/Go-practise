package main

import "fmt"

func main() {
	fmt.Println("Enter your age: ")
	var age int
	fmt.Scanln(&age)

	calAge(age)

}

func calAge(num int) {
	if num >= 18 {
		fmt.Println("Appropriate age.\nEnjoy your Movie !!")
	} else {
		fmt.Println("Movie is 18+.\nCheck for Movies within your age range")
	}
}
