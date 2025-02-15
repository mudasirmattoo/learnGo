package main

import (
	"fmt"
	"time"
)

/*
The select statement in Go is like a switch but for channels.
It allows a goroutine to wait on multiple channel operations and executes the first one that is ready.
This is useful for handling multiple concurrent communication scenarios.

select waits for one of the cases to proceed.
If multiple cases are ready, it picks one randomly.
If none are ready, the default case (if provided) executes immediately.
*/

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Start two goroutines sending data with different delays

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "hello from channel 1"
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- "hello from channel 2"
	}()

	// Use select to wait for any channel to send data
	//receive
	select {
	case msg1 := <-ch1:
		fmt.Println("Recieved : ", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recieved : ", msg2)
		//Recieved :  hello from channel 2
	}

	/*
		Two goroutines send messages to ch1 and ch2 after different delays.
		The select waits for the first available channel.
		Since ch2 has a shorter delay, select picks ch2 first and prints its message.
	*/

	useSelect()
}
