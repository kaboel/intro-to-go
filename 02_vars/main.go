package main

import (
	"fmt"
)

//var name = "Faiq" // global variable

func main() {
	/**
	Types:
	- string
	- bool
	- int int8 int16 int32 int64
	- uint uint8 uint16 int32 uint64
	- byte > alias for uint8
	- rune > alias for int32
	- float32 float64
	- complex64 complex128
	*/

	//var name = "Faiq"
	var age = 21
	var isCool = true
	//const isCool = true // constant variable
	isCool = false

	// shorthand variable declaration *only works within a function
	name := "Faiq"
	size := 41

	// multiple declaration with shorthand
	tall, email := 168, "faiq.kaboel@gmail.com"

	fmt.Println(name, age, isCool, size, tall, email)

	fmt.Printf("%T\n", name) // print "name" type
}
