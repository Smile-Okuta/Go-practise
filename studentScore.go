
package main

import "fmt"


func main() {

	fmt.Print("Enter the student's score: ")
	var score int
	fmt.Scanln(&score)

	grade := studentScore(score)

	
	fmt.Printf("The student's letter grade is: %s\n", grade)
}

func studentScore(score int) string {
	switch {
	case score >= 90 && score <= 100:
		return "A"
	case score >= 80 && score < 90:
		return "B"
	case score >= 70 && score < 80:
		return "C"
	case score >= 60 && score < 70:
		return "D"
	case score >= 0 && score <= 59:
		return "F"
	default:
		return "Please enter a score between 0 and 100."
	}
}
