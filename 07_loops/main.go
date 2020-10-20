package main

import "fmt"

func main() {
	// long way
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short way
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz test
	for i := 0; i < 100; i++ {
		if i%15 == 0 {
			fmt.Println("FizzBuzz!!!")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}
