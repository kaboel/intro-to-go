package main

import "fmt"

func main() {
	// array
	//var fruits [2]string
	//
	// assign values
	//fruits[0] = "Apple"
	//fruits[1] = "Orange"
	//
	// declare + assign
	//fruits := [2]string{
	//	"Apple",
	//	"Orange",
	//}

	// slice
	fruits := []string{
		"Apple",
		"Orange",
		"Grape",
		"Cherry",
	}

	fmt.Println(fruits)
	fmt.Println(fruits[2])
	fmt.Println(len(fruits))
	fmt.Println(fruits[0:2])
}
