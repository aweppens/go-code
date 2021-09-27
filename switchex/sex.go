package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	magniture()
}

func magniture() {

	a := os.Args
	if len(a) != 2 {
		fmt.Printf("Give me the magniture of the earthquaqe\n")
		return
	}

	r, err := strconv.ParseFloat(a[1], 64)

	if err != nil {
		fmt.Printf("I couldn't read that, sorry\n")
	}

	switch {
	case r < 2.0:
		fmt.Printf("%f is micro\n")
	case r >= 2 && r < 3:
		fmt.Printf("sfsfs")
	}
}
