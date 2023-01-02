package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number:=")
	fmt.Scanln(&n)

	fmt.Println("The sum is", n*(n+1)/2)
}
