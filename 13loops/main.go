package main

import "fmt"

func main() {

	// days := []int{3, 5, 7, 1, 46, 68, 5}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {  //you can replace _ with index
	// 	fmt.Printf("value is %v\n ", day)
	// }

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 2 {
			goto x
		}
		if rogueValue == 5 {
			// break
			rogueValue++
			continue //will skip 5
		}
		fmt.Println("value is :", rogueValue)
		rogueValue++
	}

x:
	fmt.Println("hello there")
}
