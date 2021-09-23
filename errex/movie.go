package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	oddEven()
}

func movie() {

	if a := os.Args; len(a) != 2 {
		fmt.Printf("Requires age\n")
	} else if n, err := strconv.Atoi(a[1]); err != nil || n < 0 {
		fmt.Printf("Cannot convert %q\n. Requires an age\n", a[1])
		return
	} else if n > 17 {
		fmt.Printf("R-Rated\n")
	} else if n >= 13 && n < 17 {
		fmt.Printf("PG-13\n")
	} else {
		fmt.Printf("PG-Rate\n")
	}
}

func oddEven() {

	if arg := os.Args; len(arg) != 2 {
		fmt.Printf("Pick a number\n")
	} else if n, err := strconv.Atoi(arg[1]); err != nil {
		fmt.Printf("%q is not a number\n", arg[1])
	} else if n%8 == 0 {
		fmt.Printf("%d is an even number and devisible by 8\n", n)
	} else if n%2 == 0 {
		fmt.Printf("%d is an even number\n", n)
	} else {
		fmt.Printf("%d is an odd number\n", n)
	}
}
