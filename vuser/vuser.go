package main

import (
	"fmt"
	"os"
)

func main() {

	//Program must validate username and password
	//Get input from command line
	args := os.Args

	const (
		errUn = "Invalid username %q\n"
		errPw = "Access Denied %q\n"
		msgOK = "Access Granted %q\n"

		un1, un2 = "aweppens", "ilze"
		pw1, pw2 = "12345", "qwerty"
	)

	if len(args) != 3 {
		fmt.Println("Usage: vuer [usename] [password]")
		return
	}

	u, p := args[1], args[2]
	if u != un1 && u != un2 {
		fmt.Printf(errUn, u)
	} else if u == un1 && p == pw1 || u == un2 && p == pw2 {
		fmt.Printf(msgOK, u)
	} else {
		fmt.Printf(errPw, u)
	}
}
