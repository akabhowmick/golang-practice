package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var target = rand.IntN(100) + 1

	var guess int
	attempts := 0
	const maxAttempts = 10

	for {
		fmt.Print("Guess a number between 1 and 100: ")
		fmt.Scanf("%d", &guess)

		attempts++

		if guess < target {
			fmt.Println("Too low")
		} else if guess > target {
			fmt.Println("Too high")
		} else {
			break
		}

		if attempts == 10 {
			break
		}

		if attempts != maxAttempts && attempts != 10 {
			continue
		} else {
			break
		}
	}
	fmt.Printf("\nTarget was %v ", target)
	fmt.Printf("\nYou Guessed it in %v ", attempts)
}
