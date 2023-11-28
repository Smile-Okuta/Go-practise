package main

import (
	"fmt"
)

func main() {

	var (
		speed    = 100800   // km/h
		distance = 96300000 //km
	)
	const day = 24

	var total = (distance / speed) / day

	fmt.Printf("It will take %v days to reach Mars", total)

}
