package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){

    //Loop through the program 4 times
    for x:=0; x < 4; x++ {

        //Get input from the user
        fmt.Println("Please enter a value to calulate the square root: ")
        r := bufio.NewReader(os.Stdin)
        i, err := r.ReadString('\n')
        if err != nil {
            log.Fatal("Please enter a number e.g. 2.2 ",err)
        }
        //Trim the space other wise you revceive an error
        i = strings.TrimSpace(i)

        //Conver user input from string to float
        n, err := strconv.ParseFloat(i, 64)
        if err != nil {
            log.Fatal("Please ebter a bumber e.g. 2.2 ",err)
        }

        //Calculate the square
        s, err := sqrt(n)
        if err != nil {
            log.Fatal("err3 ",err)
        }

        //Print the result
        fmt.Printf("Square root of n is %0.2f\n", s)
    }
}

func sqrt(n float64) (float64, error){
    if n < 0 {
        return 0, fmt.Errorf("number of %0.2f shoudld be greater than 0", n)
    }
    return n*n, nil
}
