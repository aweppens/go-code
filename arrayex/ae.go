package main

import (
	"fmt"
	"strings"
)

func main() {
	//exOne()
	//exTwo()
	//exTwoB()
	//exFix()
	//compareEx()
	assignEx()
}

func assignEx() {
	books := [...]string{"Dune", "Stay Golden", "James Bond"}
	upper, lower := books, books

	for i := range books {
		upper[i] = strings.ToUpper(upper[i])
		lower[i] = strings.ToLower(lower[i])
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}

func compareEx() {
	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	//Use : uint8 for one of the arrays instead of byte here.
	//Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}

func exFix() {
	names := [3]string{"Einstein", "Shepard", "Tesla"}
	books := [5]string{"Kafka's Revenge", "Stay Golden"}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}

func exTwoB() {

	names := [...]string{"Rich", "Ange", "Han"}
	distances := [...]int{50, 40, 75, 30, 125}
	data := [...]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [...]float64{3.1234}
	alives := [...]bool{true, false, true, false}
	var zero [0]byte

	_ = zero

	seperator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("\nnames", seperator)
	for i, v := range names {
		fmt.Printf("Names[%d]: %q\n", i, v)
	}

	fmt.Print("\nDistances", seperator)
	for i, v := range distances {
		fmt.Printf("Distances[%d]: %d\n", i, v)
	}

	fmt.Print("\nData", seperator)
	for i, v := range data {
		fmt.Printf("Ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nAlives", seperator)
	for i, v := range alives {
		fmt.Printf("Alives[%d]: %t\n", i, v)
	}

	fmt.Print("\nRatios", seperator)
	for i, v := range ratios {
		fmt.Printf("Ratios[%d]: %.2f\n", i, v)
	}

	fmt.Print("\nZero", seperator)
	for i, v := range zero {
		fmt.Printf("Zero[%d]: %d\n", i, v)
	}
}

func exTwo() {

	names := [3]string{"Rich", "Ange", "Han"}
	distances := [5]int{50, 40, 75, 30, 125}
	data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	ratios := [1]float64{3.1234}
	alives := [4]bool{true, false, true, false}
	var zero [0]byte

	_ = zero

	seperator := "\n" + strings.Repeat("=", 20) + "\n"

	fmt.Print("\nnames", seperator)
	for i := 0; i < len(names); i++ {
		fmt.Printf("Names[%d]: %q\n", i, names[i])
	}

	fmt.Print("\nDistances", seperator)
	for i := 0; i < len(distances); i++ {
		fmt.Printf("Distances[%d]: %d\n", i, distances[i])
	}

	fmt.Print("\nData", seperator)
	for i := 0; i < len(data); i++ {
		fmt.Printf("Ratios[%d]: %.2f\n", i, data[i])
	}

	fmt.Print("\nAlives", seperator)
	for i := 0; i < len(alives); i++ {
		fmt.Printf("Alives[%d]: %t\n", i, alives[i])
	}

	fmt.Print("\nRatios", seperator)
	for i := 0; i < len(ratios); i++ {
		fmt.Printf("Ratios[%d]: %.2f\n", i, ratios[i])
	}

	fmt.Print("\nZero", seperator)
	for i := 0; i < len(zero); i++ {
		fmt.Printf("Zero[%d]: %d\n", i, zero[i])
	}
}

func exOne() {
	var (
		names     [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float64
		alives    [4]bool
		zero      [0]byte
	)

	//Declare and print the arrays with their names
	fmt.Printf("names      :%v\n", names)
	fmt.Printf("distances  :%v\n", distances)
	fmt.Printf("data       :%v\n", data)
	fmt.Printf("ratios     :%v\n", ratios)
	fmt.Printf("alives     :%v\n", alives)
	fmt.Printf("zero       :%v\n", zero)

	//Print the types
	fmt.Printf("names      :%T\n", names)
	fmt.Printf("distances  :%T\n", distances)
	fmt.Printf("data       :%T\n", data)
	fmt.Printf("ratios     :%T\n", ratios)
	fmt.Printf("alives     :%T\n", alives)
	fmt.Printf("zero       :%T\n", zero)

	//Print the elements of the arrays
	fmt.Printf("names      :%q\n", names)
	fmt.Printf("distances  :%d\n", distances)
	fmt.Printf("data       :%d\n", data)
	fmt.Printf("ratios     :%.2f\n", ratios)
	fmt.Printf("alives     :%t\n", alives)
	fmt.Printf("zero       :%d\n", zero)
}
