package main

import (
	"fmt"
	"math/rand"
	"time"
)

//will need a funtion that will just do addition
//Test comment
func add(x, y int) int {
	return x + y
}

//Will need a funtion that will return a random nu ber
func get_rand() int {
	rand.Seed(time.Now().UnixNano())
	min := 5
	max := 50
	return rand.Intn(max-min+1) + min
}

//Will need a timer that decrases each time you
//get the answer right

// Will need the difficulty to start increasing as well
// as the score gets higher

// Will need some kind of scoring system
//Make the scoring permanents for tracking and progress

func subtractor(score int) {

}

//Function that displays the NUmbers and adds them
func adder(score int) {
	for {
		//get random numbers and set variables
		var r1 = get_rand()
		var r2 = get_rand()
		var right_answer int
		var my_answer int = 0
		//Display random numbers and calculate the correct answer
		fmt.Println("-----------------------------------------")
		fmt.Println("ADddD THE FOLLOWING")
		fmt.Println(r1)
		fmt.Println(r2)
		right_answer = add(r1, r2)

		//Get answer and display correct answer
		fmt.Println("What is the Awnser: ")
		fmt.Scanln(&my_answer)
		fmt.Println("-----")
		fmt.Println("Correct Answer: ", right_answer)
		fmt.Println("-----")
		//do the scoring
		if my_answer == right_answer {
			score++
			fmt.Println("CORRECT -- score: ", score)
			fmt.Println(" ")
		} else {
			score--
			fmt.Println("INCORRECT -- score: ", score)
			fmt.Println(" ")
		}
		if my_answer == 1010 {
			break
		}
	}
}

func main() {
	//the loop will need to include r1 and r2
	//Loop is dependend on my answer
	var score int = 0

	adder(score)
}
