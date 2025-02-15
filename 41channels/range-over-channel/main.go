package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// close(c)  -->without closing the channel, recieve func would be waiting for more values-->deadlock
		/*
			0
			1
			2
			3
			4
			fatal error: all goroutines are asleep - deadlock!

			goroutine 1 [chan receive]:
			main.main()
			        D:/Go/41channels/range-over-channel/main.go:17 +0xc5
			exit status 2
		*/

		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
