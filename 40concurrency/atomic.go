package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func implementAtomic() {

	const gr = 100

	var wg sync.WaitGroup
	wg.Add(gr)

	var counter int64

	for i := 0; i < gr; i++ {

		go func() {
			atomic.AddInt64(&counter, 1) //func atomic.AddInt64(addr *int64, delta int64) (new int64)
			runtime.Gosched()
			fmt.Println("counter\t", atomic.LoadInt64(&counter)) //func atomic.LoadInt64(addr *int64) (val int64)
			wg.Done()
		}()

		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	fmt.Println("Counter : ", counter)
}
