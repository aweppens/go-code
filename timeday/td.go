package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().Hour()
	fmt.Printf("%d\n", t)

	switch t := time.Now().Hour(); {
	case t >= 18:
		fmt.Printf("Good Evening!\n")
	case t >= 12:
		fmt.Printf("Good Afternoon\n")
	case t >= 6:
		fmt.Printf("Good Night\n")
	default:
		fmt.Printf("Good Night\n")
	}
}
