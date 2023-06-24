package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input "

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("enter the number : ")

	//comma ok syntax || error ok syntax
	

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for the inputs, ", input)
	fmt.Printf("datatype of entered number : %T \n ", input)

}
