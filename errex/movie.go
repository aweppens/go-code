package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
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
