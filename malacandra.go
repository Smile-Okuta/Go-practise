package main

import "fmt"

func main() {
	var (
		days     = 28
		distance = 56000 // km
	)
	const hour = 24

	var travel = (distance / days) / hour

	fmt.Println(travel, "km/h")

}
