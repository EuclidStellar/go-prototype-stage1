package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switch cases in Golang")
	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("value of dice is : ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("hey your dice value is 1 and you can open ")
	case 2:
		fmt.Println("hey your dice value is 2 and you can move 2 places ")
	case 3:
		fmt.Println("hey your dice value is 3 and you can move 3 places ")
	case 4:
		fmt.Println("hey your dice value is 4 and you move 4 places ")
	case 5:
		fmt.Println("hey your dice value is 5 and you move 5 places ")
	case 6:
		fmt.Println("hey your dice value is 6 and you can open or move 6 places ")
	default:
		fmt.Println("what was that ??")
	}

}
