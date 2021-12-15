package main

import "fmt"

func main (){
    a := 6
    //pass a pointer instead of a var value
    //So passing the address of the variable so that
    //The change is made there and not changing a copy
    double(&a)
    fmt.Println(&a)
    fmt.Println(a)

    truth := true
    negative(&truth)
    fmt.Println(truth)

    lies := false
    negative(&lies)
    fmt.Println(lies)
}

//Accept a pointer instead of and int value
// * means value ath the pointer address which is 6
//So update he value at the address that is passed in
func double(n *int) {
    //Update the value at the pointer
    *n *= 2
}


func negative(mybool *bool) {
    *mybool = !*mybool
}
