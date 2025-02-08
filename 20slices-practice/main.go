package main

import "fmt"

func main() {
	xs := []string{"apple", "peach", "lasan", "mango", "banana"}
	fmt.Println(xs)

	xs = append(xs[:3])
	fmt.Println(xs)

	for i, val := range xs {
		fmt.Printf("index %v and value %v\n", i, val)
	}
}
