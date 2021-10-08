package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 10

func main() {
	//	luckyNumber()
	luckyGuess()
}

func randomize() {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 11; i++ {
		n := rand.Intn(11)
		fmt.Printf("%d\n", n)
	}
}

func luckyNumber() {

	rand.Seed(time.Now().UnixNano())
	guess, i := 10, 0

	for n := 0; n != guess; {
		i++
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}

	fmt.Println()
	fmt.Printf("It took %d guesses\n", i)
	fmt.Println()
}

func luckyGuess() {

	rand.Seed(time.Now().UnixNano())
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Pick a number.")
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Be positibe!")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("YOU WIN !!!")
			return
		}
	}
	fmt.Println("YOU LSOT... Try Again?")
}
