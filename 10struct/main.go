package main

import "fmt"

func main() {

	fmt.Println("welcome to struct in Golang ")
	// no inheritance exist in golang

	gaurav := user{"gaurav", "euclidstellar@gmail.com", true, 18}
	fmt.Println(gaurav)
	fmt.Printf("guarav details are : %+v \n", gaurav)
}

type user struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
