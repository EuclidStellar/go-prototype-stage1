package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	
	fmt.Println("welcome to our Go - Hotel ")
	fmt.Println("rate my project betwenn 1 to 5 ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the rating : ", input)

	// manipulating rating

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		
		fmt.Println("added 1 to the rating entered by user ", numrating+1)

		
	}

}
