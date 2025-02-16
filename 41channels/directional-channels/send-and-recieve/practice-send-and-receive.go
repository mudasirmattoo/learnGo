package main

import (
	"fmt"
)

func mainn() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
func receivee(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
