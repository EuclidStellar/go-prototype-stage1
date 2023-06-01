package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices in golang ")
	var fruits = []string{"apple ", "potato ", "peach "}

	fmt.Printf("type of fruits list is %T \n: ", fruits)
	fruits = append(fruits, "mango ", "banana ")
	fmt.Println(fruits)

	fruits = append(fruits[1:])
	fmt.Println("this is a fruit list ", fruits)
	fmt.Println("this is a fruit list ", len(fruits))

	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 365
	highscore[2] = 456
	highscore[3] = 212

	highscore = append(highscore, 555, 767, 897)

	fmt.Println(highscore)

	sort.Ints(highscore)
	fmt.Println(highscore)

	//how tor remove a value from slice based on index

	var courses = []string{"rust", "golang", "ruby", "c++", "java "}
	//fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
