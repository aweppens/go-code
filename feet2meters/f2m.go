package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: add an argument to convert")
		return
	}

	arg := os.Args[1]
	m := .0
	f, err := strconv.ParseFloat(arg, 64)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	m = f * 0.3048
	fmt.Printf("%2.f feet is %f meters \n", f, m)

}
