package main

import "fmt"

func main() {

	fmt.Println("welcome to array in golang ")

	var fruits [5]string

	fruits[0] = "apple"
	fruits[1] = "mango"
	fruits[2] = "boobies"
	fruits[3] = "titties"
	fruits[4] = "watermelon"
	fmt.Println("your desired fruits are : ", fruits)
	fmt.Println("your desired fruits are : ", len(fruits))

	var veglist = [3]string{"potato", "mushroom", "beans "}
	fmt.Println("your desired veggies are : ", veglist)
	fmt.Println("your desired veggies are : ", len(veglist))

}
