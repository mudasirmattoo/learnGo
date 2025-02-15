package main

import "fmt"

func sendOnly() {

	ch := make(chan<- int) //creating a send-only integer channel

	go func() {
		ch <- 1
	}()

	fmt.Println(<-ch) //invalid operation: cannot receive from send-only channel ch (variable of type chan<- int)

}
