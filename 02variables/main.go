package main

import "fmt"

// global function needs to be in capital letter starting
const LoginToken string = "BLAHHHHH BLAHHHH" //PUBLIC

func main() {
	fmt.Println("variables")
	var username string = "gaurav"
	fmt.Println(username)
	fmt.Printf("variable is of datatype :%T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of datatype : %T \n", isLoggedIn)

	var smallfloat float32 = 256.8685766
	fmt.Println(smallfloat)
	fmt.Printf("variable is of datatype : %T \n", smallfloat)

	var smallfloat2 float64 = 256.868576656435
	fmt.Println(smallfloat2)
	fmt.Printf("variable is of datatype : %T \n", smallfloat2)

	//default values and some aliases

	var defaultvalue int
	fmt.Println(defaultvalue)
	fmt.Printf("variable is of datatype : %T \n", defaultvalue)

	//implicit type

	var website = "http:://euclidstellar.vercel.app"
	fmt.Println(website)
	fmt.Printf("variable is of datatype : %T \n", website)

	//no var style

	numberofuser := 300000
	fmt.Println(numberofuser)
	fmt.Printf("variable is of datatype : %T \n", numberofuser)

	//datatype 2

	numberofuser1 := 300000.88372646
	fmt.Println(numberofuser1)
	fmt.Printf("variable is of datatype : %T \n", numberofuser1)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of datatype : %T \n", LoginToken)

}
