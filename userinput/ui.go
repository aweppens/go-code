package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
    "strconv"
)

func main(){

    fmt.Println("Enter grade: ")
    reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println(err, "Pleasr enter a number")
        return
    }

    input = strings.TrimSpace(input)
    grade, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Please enter a number")
        return
    }

    if grade >= 60 {
        fmt.Println("PASS")
    } else {
        fmt.Println("FAIL")
    }
}
