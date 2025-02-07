package main

import "fmt"

func main() {

	var loginCount int
	fmt.Println("enter login count :")
	fmt.Scanln(&loginCount)

	var result string
	if loginCount <= 10 {
		result = "regular user"
	} else if loginCount > 10 && loginCount < 15 {
		result = "okayish"
	} else {
		result = "something wrong"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 11; num < 10 {
		fmt.Println("num is less than 10")
	}

}
