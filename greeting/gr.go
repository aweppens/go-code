package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("When running this program input 5 names after the command as arguments")
	fmt.Println("There are:", len(os.Args)-1, "names")
	fmt.Println("Hi", os.Args[1])
	fmt.Println("Hi", os.Args[2])
	fmt.Println("Hi", os.Args[3])
	fmt.Println("Hi", os.Args[4])
	fmt.Println("Hi", os.Args[5])

}
