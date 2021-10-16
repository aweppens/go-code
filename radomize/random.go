package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 10
const usage = `
    Welcome to the Lucky Number Game! 🍀
    The program will pick %d random numbers.
    Your mission is to guess one of those numbers.
    The greater your number is, harder it gets.
    Wanna play?

    Example:
    random 5 8
    random -v 5 8
`

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

	//Variables
	rand.Seed(time.Now().UnixNano())
	c := 0
	var msg string
	args := os.Args[1:]

	//Must be 2 numbers or -v or -h arguments
	//So there can either be 2 or 3 args
	if len(args) < 2 {
		fmt.Println(usage, maxTurns)
		return
	}

	var verbose bool

	if args[0] == "-v" {
		verbose = true
	}

	//guess, err := strconv.Atoi(args[0])
	guess, err := strconv.Atoi(args[len(args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	guess2, err := strconv.Atoi(args[len(args)-1])
	//guess2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 || guess2 < 0 {
		fmt.Println("Be positibe!")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess2 + 1)
		c++

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess && c == 1 || n == guess2 && c == 1 {
			fmt.Printf("The player player won on firt try: %d\n", guess)
			return
		}

		r := rand.Intn(3)
		switch {
		case r == 1:
			msg = "\nYou're a weaner!!"
		case r == 2:
			msg = "\nJackpot!!!"
		case r == 3:
			msg = "\nBaby, you're on fire!!!"
		}

		if n == guess || n == guess2 {
			fmt.Println(msg)
			return
		}
	}

	r := rand.Intn(3)
	switch {
	case r == 1:
		msg = "\nSorry betterluck next time."
	case r == 2:
		msg = "\nComputer says no"
	case r == 3:
		msg = "\nYou Lose.. Try again"
	}
	fmt.Println(msg)
}
