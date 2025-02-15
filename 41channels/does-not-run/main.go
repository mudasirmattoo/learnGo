package main

import "fmt"

/*
Channels
This is where channels come in. Channels are a synchronization primitive that allow Goroutines
to communicate and synchronize execution. You can send values into a channel and receive
those values into another Goroutine. This allows you to coordinate tasks and send data safely
between Goroutines:
```go
ch := make(chan int) // Make a channel of type int
go func() {
ch <- 1 // Send the value 1 into the channel
}()
fmt.Println(<-ch) // Receive the value from the channel and print it
*/
func main() {

	ch := make(chan int)

	ch <- 1 //send/load value into the channel

	fmt.Println(ch) // it will throw an error  --> deadlock occurs

	/*
		ALWAYS REMEMBER:
		CHANNEL BLOCKS,
		which means if you load a value into a channel, there should be a goroutine to receive that,
		otherwise it goes into deadlock
	*/

}
