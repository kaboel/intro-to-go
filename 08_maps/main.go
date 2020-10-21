package main

import "fmt"

func main() {
	// define map
	//emails := make(map[string]string)
	//
	//// assign key value
	//emails["bob"] = "bob@gmail.com"
	//emails["sharon"] = "sharon@gmail.com"
	//emails["mike"] = "mike@gmail.com"

	// declare map and key value
	emails := map[string]string{
		"bob":    "bob@gmail.com",
		"sharon": "sharon@gmail.com",
	}
	emails["mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["bob"])

	// delete a key
	delete(emails, "bob")
	fmt.Println(emails)
}
