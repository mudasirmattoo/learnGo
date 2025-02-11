package main

import "fmt"

func main() {
	fmt.Println(adder(5, 6))
}

func adder(a int, b int) int {
	return a + b
}
