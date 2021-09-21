package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("There are:", len(os.Args)-1, "names")
	fmt.Println("Hi", os.Args[1])
	fmt.Println("Hi", os.Args[2])
	fmt.Println("Hi", os.Args[3])
	fmt.Println("Hi", os.Args[4])
	fmt.Println("Hi", os.Args[5])

}
