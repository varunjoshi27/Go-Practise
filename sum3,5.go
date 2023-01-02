package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter the number:=")
	fmt.Scanln(&num)
	if (num%3 == 0) || (num%5 == 0) {
		fmt.Println("The sum is ", num*(num+1)/2)
	} else {
		fmt.Println("Unrecognised number")
	}
}
