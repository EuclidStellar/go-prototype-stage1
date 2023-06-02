package main

import "fmt"

func main() {
	fmt.Println("welcome to if else in golang ")
	logincount := 23
	var result string
	if logincount < 10 {
		result = "regular user"
	} else if logincount > 10 {
		result = "watchout"

	} else {
		result = "exactly 10 logins "
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("number is even ")
	} else {
		fmt.Println("number is odd ")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10 ")
	} else {
		fmt.Println("num is not less than 10 ")
	}

}
