package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//	age()
	//	bigSphere()
	//giveMeArgs()
	vowelOrConsonant()
}

func vowelOrConsonant() {

	var (
		args = os.Args
		a    = len(args) - 1
	)

	//check if there are args
	if a == 0 {
		fmt.Println("Give me lines")
		//if there are one or more args check the lenght
	} else if a >= 1 {
		l := len(args[1])
		if l > 1 {
			fmt.Println("Give me a letter")
		} else if l == 1 {
			s := strings.IndexAny(args[1], "aeiou")
			y := strings.IndexAny(args[1], "yw")
			if s == 0 {
				fmt.Printf("%s is a vowel\n", args[1])
			} else if y == 0 {
				fmt.Printf("%s is ometimes a vowel, sometimes not\n", args[1])
			} else {
				fmt.Printf("%s is a consonant\n", args[1])
			}
		}
	}
}

func giveMeArgs() {
	var (
		args = os.Args
		l    = len(args) - 1
	)

	fmt.Println(l)

	if l == 0 {
		fmt.Println("Give me lines")
	} else if l == 1 {
		fmt.Printf("There is one: %d\n", l)
	} else if l == 2 {
		fmt.Printf("There are two: %d\n", l)
	} else if l == 3 {
		fmt.Printf("There are three: %d\n", l)
	} else if l == 4 {
		fmt.Printf("There are four: %d\n", l)
	} else if l == 5 {
		fmt.Printf("There are five: %d\n", l)
	} else {
		fmt.Println("You clearly love words")
	}

}

func age() {

	age := 30

	if age > 60 {
		fmt.Printf("Getting older, you are %d!\n", age)
	} else if age > 30 {
		fmt.Printf("Getting wiser, you are %d!\n", age)
	} else if age > 20 {
		fmt.Printf("Adulthood, you are %d!\n\n", age)
	} else if age > 10 {
		fmt.Printf("Young Blood, you are %d!\n", age)
	} else {
		fmt.Printf("Booting up\n")
	}
}

func bigSphere() {

	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big Spere")
	} else {
		fmt.Println("I don't know")
	}
}
