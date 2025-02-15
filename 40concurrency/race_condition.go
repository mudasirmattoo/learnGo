package main

import (
	"fmt"
	"runtime"
	"sync"
)

func raceCondition() {

	const gr = 10 //set goroutines to 10
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {

		go func() {
			v := counter // v is local for each goroutine
			// time.Sleep(time.Second)
			// runtime.Gosched()   /runtime.Gosched() does not guarantee that all goroutines will execute in order.
			v++
			counter = v
			fmt.Println("Counter : ", counter)
			wg.Done()
		}()
		/*
			Why Does time.Sleep(time.Second) Result in Counter = 1?
			When you uncomment time.Sleep(time.Second), here's what happens:

			1. The first goroutine runs and sleeps for 1 second.
			2. Meanwhile, all 10 goroutines are scheduled and start executing.
			3. counter = v only executes after all goroutines have already read counter = 0.
			4. Each goroutine increments its local v (which is 0), and when they write back, they all overwrite counter = 1.
			5. Final counter = 1, as all goroutines wrote the same value!
			6. Goroutines count goes from 2 → 11 → 1 because they get scheduled and execute one by one.

			Why Does runtime.Gosched() Result in Counter = 10?
			When you use runtime.Gosched(), this happens:

			1. runtime.Gosched() yields the processor to allow other goroutines to run.
			2. Since all 10 goroutines are competing to increment counter, most of them successfully increment it.
			3. The counter reaches 10 because scheduling allows multiple goroutines to execute in a somewhat orderly manner.
			4. Goroutine count stays at 2 throughout because goroutines don’t pile up like in the time.Sleep case.
		*/

		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	fmt.Println("Counter end value: ", counter)
}
