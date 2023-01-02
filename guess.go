package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	min, max := 1, 10
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(max-min) + min

	var number int

	for chances := 2; chances >= 0; chances-- {
		fmt.Println("Enter the number between 0 to 10:=. Chances left", (chances + 1))
		fmt.Scanln(&number)
		if (number > 10) || (number < 0) {
			fmt.Println("Please play carefully")
			os.Exit(0)
		}

		if number == secretNumber {
			fmt.Printf("You won the match")
		} else if chances > 0 {
			fmt.Printf("Wrong guess, please try again")
		} else {
			fmt.Println("you have lost. play better!")
		}
	}

}
