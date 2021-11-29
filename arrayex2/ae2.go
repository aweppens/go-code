package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
    //rates()
    //searchBooks()
    average()
}

func average(){

    args := os.Args[1:]
    if len(args) < 1 || len(args) > 5 {
        fmt.Println("Please tell me numbers (max 5 numbers)")
        return
    }
}

func searchBooks(){

    books := [...]string{
       "Kafka's Revenge",
       "Stay Golden",
       "Hitchhikers Guide",
       "Kafka's Revenge 2nd Edition",
    }

    //name := "Kafka"
    args := os.Args[1:]
    if len(args) != 1 {
        fmt.Println("Please enter a title to search for.")
        return
    }

    name := strings.ToLower(args[0])

    fmt.Println("Search results: ")
    var found bool
    for _, i := range books {
        //fmt.Println(books[i])
        if strings.Contains(strings.ToLower(i), name) {
            fmt.Printf("+ %s\n", i)
            found = true
        }
    }

    if !found {
        fmt.Printf("We don't have the book: %q\n", name)
    }
}


func rates(){
    const (
        EUR = iota
        GBP
        JPY
    )

    rates := [...]float64{
        EUR: 0.88,
        GBP: 0.78,
        JPY: 113.02,
    }

    args := os.Args[1:]
    if len(args) != 1{
        fmt.Println("Please provide the amoun to be converted.")
        return
    }

    amount, err := strconv.ParseFloat(args[0], 64)
    if err != nil {
        fmt.Println("Invalid amount. It should be a number.")
        return
    }

    fmt.Printf("%.2f USD is %.2f EUR\n", amount, rates[EUR]*amount)
    fmt.Printf("%.2f USD is %.2f GBP\n", amount, rates[GBP]*amount)
    fmt.Printf("%.2f USD is %.2f JPY\n", amount, rates[JPY]*amount)
}

func scientists(){
    fmt.Println("Program starting...")

    names := [...][3]string{
        {"First Name", "Last Name", "Nick Name"},
        {"Albert", "Einstein", "emc2"},
        {"Issac", "Newton", "Apple"},
        {"Stephen", "Hawking", "Blackhole"},
        {"Marie", "Curie", "Raddium"},
        {"Charles", "Darwin", "Fittest"},
    }

    for i := range names {
        n := names[i]
        fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

        if i == 0{
            fmt.Println(strings.Repeat("=", 50))
        }
    }
}

