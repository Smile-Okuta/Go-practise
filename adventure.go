package main

import "fmt"

func main() {
	fmt.Println("What room are you at? ")
	var room string
	fmt.Scanln(&room)

	adventureRoom(room)

}

func adventureRoom(name string) {
	if name == "cave" || name == "Cave" {
		fmt.Println("You find yourself in a dimly lit cavern")
	} else if name == "entrance" || name == "Entrance" {
		fmt.Println("There is a cavern entrance here and a path to the east")
	} else if name == "mountain" || name == "Mountain" {
		fmt.Println("There is a clif here. A path leads west down the mountain")
	} else {
		fmt.Println("Everything is white")
	}

}
