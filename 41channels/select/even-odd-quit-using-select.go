package main

import "fmt"

func useSelect() {

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	recieve(even, odd, quit)

	fmt.Println("about to exit")

}

func send(e, o, q chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0 //Closing quit is unnecessary because recieve terminates after reading from it.
}

func recieve(e, o, q <-chan int) {
	for { //receive till return statement
		select {
		case v := <-e: //received through even channel
			fmt.Println(v, " Received through even channel")
		case v := <-o: //received through odd channel
			fmt.Println(v, " Received through odd channel")
		case v := <-q: //received through quit channel
			fmt.Println(v, " Received through quit channel")
			return
		}
	}
}

/*
OUTPUT
----------------
Recieved :  hello from channel 2
0  Received through even channel
1  Received through odd channel
2  Received through even channel
3  Received through odd channel
4  Received through even channel
5  Received through odd channel
6  Received through even channel
7  Received through odd channel
8  Received through even channel
9  Received through odd channel
10  Received through even channel
11  Received through odd channel
12  Received through even channel
13  Received through odd channel
14  Received through even channel
15  Received through odd channel
16  Received through even channel
17  Received through odd channel
18  Received through even channel
19  Received through odd channel
0  Received through quit channel
about to exit
*/
