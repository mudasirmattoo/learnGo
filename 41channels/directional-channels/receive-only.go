package main

import "fmt"

func receiveOnly() {

	ch := make(<-chan int) //creating a recieve-only integer channel

	go func() {
		ch <- 1 //invalid operation: cannot send to receive-only channel ch (variable of type <-chan int)
	}()

	fmt.Println(<-ch)

}
