package main

import (
	"fmt"
	//"sync"
)

func main() {

	c := make(chan int)
	// var wg sync.WaitGroup
	//wg.Add(2)
	done := make(chan bool) //make semaphore channel

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// wg.Done()
		done <- true
	}()
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// wg.Done()
		done <- true
	}()

	go func() {
		// wg.Wait()
		<-done //take the value of the done channel
		<-done
		close(c)
	}()

	/*
		WRONG WAY TO IMPLEMENT SEMAPHORE
		-------------------------------

		<-done  // when one goroutine puts value into c channel, there should be someone to receive, but using <-done, we never reach to the below code
		<-done
		close(c)
	*/

	for i := range c {
		fmt.Println(i)
	}

}
