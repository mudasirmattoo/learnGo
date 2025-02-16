package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func receive(c, q <-chan int) {

	select {
	case int1 := <-c:
		fmt.Println("received from channel c1", int1)
	case int2 := <-q:
		fmt.Println("received from channel c1", int2)
	}
}
