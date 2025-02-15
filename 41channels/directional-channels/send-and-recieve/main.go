package main

import "fmt"

func main() {
	// Create an unbuffered channel for ints
	c := make(chan int)

	// Start a goroutine that sends a value into the channel
	go send(c)

	// Call receive synchronously (main func waits here until a value is received)
	receive(c)

	//if we use go with recieve, it might not be able to recieve the value before main func gets terminated
	fmt.Println("about to exit")
}

func send(ch chan<- int) {
	// Send the value 33 into the channel (this blocks until a receiver is ready)
	ch <- 33
}

func receive(ch <-chan int) {
	fmt.Println(ch) //0xc000062070(any address) if we use only ch
	// Read a value from the channel (this blocks until a value is available)
	fmt.Println(<-ch) //33
}

/*
REASON FOR NOT USING GO WITH THE RECIEVE() FUNCTION AS WELL:
-------------------------------------------------------------
Both send(c) and receive(c) will run concurrently in separate goroutines.
The main function would start both goroutines and then immediately exit (unless you add some synchronization, like a WaitGroup or sleep). When the main function returns, the program terminates, often before the receive goroutine has a chance to read and print the value.
This can result in incomplete output or no output at all.

*/
