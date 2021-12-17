package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "log"
)

//Create a getfloats function to read filename and return []float64 and err
func getFloats(filename string)([]float64, error){

    //Variable that stores []numbers, will contain nil by default... Append treats it like an empty slice
    var numbers []float64

    //Open the filename
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }

    //Read the file line by line and return the [] and err
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //Read each number
        number, err := strconv.ParseFloat(scanner.Text(), 64)
        if err != nil {
            return nil, err
        }
        //append each number to the slice
        numbers = append(numbers, number)
    }

    //Close file
    err = file.Close()
    if err != nil {
        return nil, err
    }

    if scanner.Err() != nil {
        return nil, scanner.Err()
    }

    return numbers, nil
}

//Create the main function which will get data and work out the average of the data
func main(){

    arg := os.Args[1:]
    if len(arg) != 1 {
        fmt.Println("Please enter the the name of the data file. Only 1 argument allowed")
        return
    }

    //Get the data from the file by passing file to getfloats function and error Check
    numbers, err := getFloats(arg[0])
    if err != nil {
        log.Fatal(err)
    }

    //create a var to hold the sum
    var sum float64 = 0

    //Loop throught the array and and add the values
    for _, n := range numbers {
        sum += n
    }

    //Get the number of datapoints of the slice using len()
    datapoints := float64(len(numbers))

    //Print the avaragle
    fmt.Printf("The avarage of sum %0.2f is %0.2f\n", sum, sum/datapoints)

}

