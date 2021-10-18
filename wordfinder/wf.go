package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const corpus = "lazy cat jumps again and again and agian"

func main() {
	//wordFinder()
	pathFinder()
}

func wordFinder() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

	//Loops though the user input
	for _, q := range query {
		//Program shold be case incensitive
		q = strings.ToLower(q)
		//Loops though the corpus words
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				//Stops when finding a mach to avoid printing duplicates (unique matches only)
				break
			}
		}
	}
}

func pathFinder() {
	path := filepath.SplitList(os.Getenv("PATH"))
	fmt.Println(path)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	query := os.Args[1:]

	for _, q := range query {
		for i, p := range path {
			if q == p {
				fmt.Printf("#%-2d: %q\n", i+1, p)
			}
		}
	}
}
