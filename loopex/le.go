package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//mtable()
	sum()
}

func sum() {

	args := os.Args

	if len(args) != 3 {
		fmt.Printf("Need 2 arguments:min and max values\n")
		return
	}

	min, maxerr := strconv.Atoi(args[1])
	max, minerr := strconv.Atoi(args[2])
	if maxerr != nil || minerr != nil {
		fmt.Printf("I need a numbers\n")
		return
	}

	var sum int
	for i := min; i <= max; i++ {

		if i%2 != 0 {
			continue
		}

		sum += i
		fmt.Print(i)

		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

func mtable() {

	max := os.Args

	if len(max) != 2 {
		fmt.Printf("Need an argument for the dyamic table\n")
		return
	}

	m, err := strconv.Atoi(max[1])
	if err != nil {
		fmt.Printf("I need a number\n")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= m; i++ {
		fmt.Printf("%5d", i)
	}

	fmt.Println()

	for i := 0; i <= m; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= m; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}

}
