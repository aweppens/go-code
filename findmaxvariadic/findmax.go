package main

import (
    "fmt"
    "math"
)

func main(){

    fmt.Printf("Max no for A is %0.2f\n", maximum(12.4, 23.5, 23.5, 243.34))
    fmt.Printf("Max no for B is %0.2f\n", maximum(12.4, 12.3, 45.6))

    fmt.Printf("Range for for A is %0.2f\n", inRange(5, 50, 1,4,6,9,10, 100))
    fmt.Printf("Range for B is %0.2f\n", inRange(20, 30, 12.3, 50.6, 40, 30, 100.54, 500.1))
}

func inRange(min float64, max float64, numbers ...float64) []float64 {

    var result []float64

    for _, n := range numbers {
        if n >= min && n <= max {
            result = append(result, n)
        }
    }
    return result
}

func maximum(numbers ...float64) float64{
    max := math.Inf(-1)

    for _, n := range numbers {
        if n > max {
            max = n
        }
    }
    return max
}
