package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// read pointer value
	fmt.Println(*b)

	// change pointer value, this means a
	*b = 7
	fmt.Println(a, *b)
}
