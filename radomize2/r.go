package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	//"strconv"
	"bufio"
	"time"
)

func main(){
    //Generate a random number and store it in r
    rand.Seed(time.Now().UnixNano())
    r := rand.Intn(100 + 1)
    fmt.Printf("Rand dome no: %d\n", r)
    fmt.Println("I've chosen a random number between 1 and 100")
    fmt.Println("Can you guess ehat it is? You have 10 guesses. Good Luck player.")
    reader := bufio.NewReader(os.Stdin)

    success := false
    for i := 0; i < 10; i++ {
        fmt.Println("Please enter your guess")
        input, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input = strings.TrimSpace(input)
        guess, err := strconv.Atoi(input)
        if err != nil {
            log.Fatal(err)
        }

        if guess < r {
            fmt.Println("TOO LOW!")
        } else if guess > r {
            fmt.Println("TOO HIGH!!!")
        } else {
            success = true
            fmt.Println("YOU WIN!!!")
            break
        }

        fmt.Printf("Your have %d gueeses left\n", 9-i)
    }

    if !success {
        fmt.Printf("Sorry NO LUCK :-( It was: %d\n", r)
    }
}
