package main

import "fmt"

func main() {

	ch := make(chan int, 2) //store 2 values in the buffer channel

	ch <- 1
	ch <- 2
	fmt.Println(<-ch) //1
	fmt.Println(<-ch) //2
	fmt.Println(<-ch) // deadlock, won't print
	/*
		1
		2
		fatal error: all goroutines are asleep - deadlock!

		goroutine 1 [chan receive]:
		main.main()
		        D:/Go/41channels/buffer-channel/main.go:13 +0x116
		exit status 2
	*/

	/*
		ALWAYS REMEMBER:
		CHANNEL BLOCKS,
		which means if you load a value into a channel, there should be a goroutine to receive that,
		otherwise it goes into deadlock
	*/

}
