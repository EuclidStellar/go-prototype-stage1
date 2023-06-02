package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to loops in golang ")
	days := []string{"sunday", "tuesday", "wednesday", "thursday", "friday"}
	fmt.Println(days)

	//Type 1

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])

	// }

	//Type 2

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	//Type 3

	for index, days := range days {
		fmt.Printf("index id %v and value is %v\n", index, days)
	}
	//Type 4
	// for _, days := range days {
	// 	fmt.Printf("index id and value is %v\n", days)
	// }

	roughvalue := 1
	for roughvalue < 10 {
		if roughvalue == 2 {
			goto gg
		}
		// if roughvalue == 5 {
		// 	roughvalue++
		// 	continue //will skip 5
		// 	//break will terminate the loop

		// }
		fmt.Println("value is : ", roughvalue)
		roughvalue++

	}
gg:
	fmt.Println("euclid stellar")

}
