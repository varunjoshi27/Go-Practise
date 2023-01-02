package main

import "fmt"

func main() {
	var person string

	fmt.Println("Enter the name:=")
	fmt.Scanln(&person)
	if (person == "Alice") || (person == "BoB") {
		fmt.Println("Hello", person)
	} else {
		fmt.Println("User not recognised")
	}
}
