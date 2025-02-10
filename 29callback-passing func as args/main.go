package main

import "fmt"

func main() {

	z := doMath(9, 8, adder)
	v := doMath(9, 8, subtract)
	fmt.Println(z)
	fmt.Println(v)
}

func adder(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
