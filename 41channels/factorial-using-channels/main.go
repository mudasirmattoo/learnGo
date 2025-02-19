package main

import "fmt"

func main() {
	c := factorial(5)

	for i := range c {
		fmt.Println(i)
	}
}

func factorial(n int) chan int {
	out := make(chan int)
	total := 1
	go func() {
		for i := n; i > 0; i-- {
			total *= i
		}
		out <- total
		close(out)
	}()
	return out
}
