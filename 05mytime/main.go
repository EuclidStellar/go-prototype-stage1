package main

import (
	"fmt"
	"time"

)

func main() {
	fmt.Println("welcome to time study of golang ")
	presenttime := time.Now()
	fmt.Println(presenttime)
	fmt.Println(presenttime.Format("01-04-2006  15:04:05 Monday"))

	createwddate := time.Date(2005, time.April, 25, 4, 4, 5, 8, time.UTC)
	fmt.Println(createwddate)
	fmt.Println(createwddate.Format("01-02-2006 Monday"))

}
