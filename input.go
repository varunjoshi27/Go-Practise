package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name:=")
	fmt.Scanln(&name)
	fmt.Println("Your name is ", name)
}
