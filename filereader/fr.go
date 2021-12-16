package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
)

func main() {

    file, err := os.Open("data")
    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)
    //Loops until the end of the file is reached
    //Read a line from the file
    for scanner.Scan() {
        //Print the line
        fmt.Println(scanner.Text())
    }

    //Close the file to free resources
    err = file.Close()
    //If theres an error closing the file report and close
    if err != nil {
        log.Fatal(err)
    }

    //if theres an error scanning the file report and close
    if scanner.Err() != nil {
        log.Fatal(scanner.Err())
    }
}
