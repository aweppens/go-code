package main

import (
	"fmt"
	"os"
)

func main() {

	age := os.Args[1]
	name := "Inanc"
	fmt.Printf("my age is %s years old.\n", age)
	fmt.Printf("My name is %s Gumus \n", name)

	tf := false
	fmt.Printf("These are %t claims\n", tf)
	temp := 14.34534
	fmt.Printf("The temp in SA is %.1f degrees\n", temp)

	fmt.Printf("%q\n", "Hollow World")

	i := 3
	fmt.Printf("Type of 4 is %T \n", i)

	f := 3.14
	fmt.Printf("Type of %.2f is %T \n", f, f)
}
