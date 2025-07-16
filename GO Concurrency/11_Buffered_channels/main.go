package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Println("Got", i, "from channel")

		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 50; i++ {
		fmt.Println("Sending", i, "to the channel")
		ch <- i
		fmt.Println("Sent", i, "to the channel")
	}
	fmt.Println("Done")
	close(ch)
}

// Sending 0 to the channel
// Sent 0 to the channel
// Sending 1 to the channel
// Sent 1 to the channel
// Sending 2 to the channel
// Sent 2 to the channel
// Sending 3 to the channel
// Sent 3 to the channel
// Sending 4 to the channel
// Sent 4 to the channel
// Sending 5 to the channel
// Sent 5 to the channel
// Sending 6 to the channel
// Sent 6 to the channel
// Sending 7 to the channel
// Sent 7 to the channel
// Sending 8 to the channel
// Sent 8 to the channel
// Sending 9 to the channel
// Sent 9 to the channel
// Sending 10 to the channel
// Sent 10 to the channel
// Sending 11 to the channel
// Got 0 from channel
// Got 1 from channel
// Sent 11 to the channel
// Sending 12 to the channel
// Got 2 from channel
// Sent 12 to the channel
// Sending 13 to the channel
// Got 3 from channel
// Sent 13 to the channel
// Sending 14 to the channel
// Got 4 from channel
// Sent 14 to the channel
// Sending 15 to the channel
// Got 5 from channel
// Sent 15 to the channel
