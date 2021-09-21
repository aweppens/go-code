package main

import (
	"fmt"
)

func main() {

	const (
		Winter = (iota + 1) * 3
		Spring
		Summer
		Fall
	)

	fmt.Println(Winter, Spring, Summer, Fall)
}
