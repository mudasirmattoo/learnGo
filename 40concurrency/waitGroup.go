package main

import (
	"fmt"
	"runtime"
	"sync"
)

// ????check method sets
var wg sync.WaitGroup //wg created is not a type pointer , but  func (wg *WaitGroup) Add(delta int) accepts the receiver of pointer to waitGroup interface

func useWaitGroup() {

	/*type WaitGroup
	func (wg *WaitGroup) Add(delta int)
	func (wg *WaitGroup) Done()
	func (wg *WaitGroup) Wait()*/

	wg.Add(1) //Add adds delta, which may be negative, to the [WaitGroup] counter. If the counter becomes zero, all goroutines blocked on [WaitGroup.Wait] are released. If the counter goes negative, Add panics.
	go hello1()
	world2()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() //Wait blocks until the [WaitGroup] counter is zero
}

func hello1() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello ", i)
	}
	wg.Done() //Done decrements the [WaitGroup] counter by one.
}

func world2() {
	for i := 0; i < 5; i++ {
		fmt.Println("world ", i)
	}
}
