package main

import "fmt"

func main() {

	// if else
	x := 5
	y := 10

	if x < y { // && and || operators are also work
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is more than %d\n", x, y)
	}

	// if else if
	color := "Red"

	if color == "Red" {
		fmt.Println("Color is Red")
	} else if color == "Blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color isn't Red or Blue")
	}

	// switch
	switch color {
	case "Red":
		fmt.Println("Color is Red")
	case "Blue":
		fmt.Println("Color is Blue")
	default:
		fmt.Println("Color isn't Red or Blue")
	}
}
