/*
How Fan-in Works
Multiple Channels:
You have several channels (each may be producing data concurrently).

Fan-in Function:
A fan-in function (or goroutine) listens to all these channels using a select statement, and whenever a value is received from any channel, it forwards the value to a single output channel.

Single Consumer:
The consumer reads from the single output channel, receiving data from all the original channels in a merged order.

*/

package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	//send
	go send("Sending A :", ch1)
	go send("Sending B :", ch2)

	//recieve combined channel
	merged := fanInn(ch1, ch2)

	for i := range merged {
		fmt.Println("recieved ", i)
	}

	fmt.Println("about to exit")
}

func send(name string, c chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println(name, i)
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}

// fanIn merges two channels into a single output channel.
// it will return a combined channel that recieves int values
func fanInn(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		// Keep reading until both channels are closed.
		for ch1 != nil || ch2 != nil {
			select {
			case v, ok := <-ch1: //check if channel 1 is sending any value
				if !ok { //channel 1 is closed
					ch1 = nil // Set it to nil so select won't pick it again.
					continue
				}
				out <- v //put the value into the combined channel
			case v, ok := <-ch2: //check if channel 1 is sending any value
				if !ok { //channel 1 is closed
					ch2 = nil // Set it to nil so select won't pick it again.
					continue
				}
				out <- v //put the value into the combined channel
			}
		}
		close(out) // All channels are done; close the output channel.

	}()
	return out //the merged channel
}

/*
OUTPUT
--------------------
Sending A : 0
Sending B : 0
recieved  0
recieved  0
Sending B : 1
recieved  1
Sending A : 1
recieved  1
Sending A : 2
recieved  2
Sending B : 2
recieved  2
Sending A : 3
recieved  3
Sending B : 3
recieved  3
Sending A : 4
recieved  4
Sending B : 4
recieved  4
about to exit
*/
