package main

import "fmt"

func main() {

	//setting up the pipeline
	c := gen(3, 4)  //gen() produces a channel
	square := sq(c) //consumes a channel

	for i := range square {
		fmt.Println(i)
	}

	fmt.Println("quit baby")
}

func gen(n ...int) chan int {

	out := make(chan int)
	go func() {
		for _, i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func sq(in chan int) chan int {

	out := make(chan int)
	go func() {
		for k := range in {
			out <- k * k
		}
		close(out)
	}()
	return out
}
