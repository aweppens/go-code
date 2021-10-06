package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	luckyNumber()
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

	for n := 0; n != guess; n++ {
		i += n
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}

	fmt.Printf("It took %d amount of guessesl\n", i)
	fmt.Println()
}
