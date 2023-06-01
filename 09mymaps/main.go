package main

import "fmt"

func main() {
	fmt.Println("welcome to maps of go lang ")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["cpp"] = "c"

	fmt.Println("list of all languages : ", languages)
	fmt.Println("js shorts for : ", languages["js"])
	fmt.Println("py shorts for : ", languages["py"])
	fmt.Println("cpp shorts for : ", languages["c"])

	delete(languages, "py")

	fmt.Println("list of languages : ", languages)

	//loops are intresting in golang

	// for key, value := range languages {
	// 	fmt.Printf("for key %v , value is %v \n", key, value)

	// }

	for _, value := range languages {
		fmt.Printf("for key v , value is %v \n", value)

	}

}
