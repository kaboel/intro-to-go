package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Faiq"))
	fmt.Println(getSum(3, 4))
}

func greet(name string) string {
	return "Hello" + " " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}
