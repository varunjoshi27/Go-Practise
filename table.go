package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number to print the multiplication table:")
	fmt.Scanf("%d", &n)
	for i := 1; i < 11; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}
}
