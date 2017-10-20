package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Guess the number!")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	secretNumber := r.Intn(101)
	// fmt.Printf("The secret number is %d\n", secretNumber)
	for {
		fmt.Println("Please input your guess")
		var guessStr string
		fmt.Scanln(&guessStr)
		guess, err := strconv.Atoi(guessStr)
		if err != nil {
			continue
		}
		fmt.Printf("You guessed: %d\n", guess)
		if guess > secretNumber {
			fmt.Println("Too big!")
		} else if guess < secretNumber {
			fmt.Println("Too small!")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}
