package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	// specific(cr or cs) to general doesn't assign
	//c = cr //cannot use cr (variable of type <-chan int) as chan int value in assignment
	//c = cs //cannot use cs (variable of type chan<-int) as chan int value in assignment

	// specific to specific doesn't assign
	//cs = cr  //cannot use cr (variable of type <-chan int) as chan<- int value in assignment

	// general(c) to specific(cs/cr) assigns
	cr = c
	cs = c
}
