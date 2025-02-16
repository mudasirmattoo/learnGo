package main

import "fmt"

func main() {
	// Create a slice of 10 channels, each of type chan int.
	channels := make([]chan int, 10)

	// Initialize each channel in the slice.
	for i := range channels {
		channels[i] = make(chan int)
	}

	// Launch goroutines that send values on each channel.
	for i, ch := range channels {
		go func(i int, c chan int) {
			c <- i
		}(i, ch)
	}

	// Receive from each channel.
	for i, ch := range channels {
		val := <-ch
		fmt.Printf("Channel %d received: %d\n", i, val)
	}

	fmt.Println("all done")
}
