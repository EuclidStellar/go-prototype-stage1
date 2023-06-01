package main

import "fmt"

func main() {

	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("value of pointer is : ", ptr)

	mynum := 23
	var ptr = &mynum
	fmt.Println("value of actual pointer is : ", ptr)  // address of value ( what's inside the pointer )
	fmt.Println("value of actual pointer is : ", *ptr) // gives value

	*ptr = *ptr + 2
	fmt.Println("New value is : ", mynum)

}
