package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
type RWMutex
func (rw *RWMutex) Lock()
func (rw *RWMutex) RLock()  //lock reading
func (rw *RWMutex) RLocker() Locker
func (rw *RWMutex) RUnlock()
func (rw *RWMutex) TryLock() bool
func (rw *RWMutex) TryRLock() bool
func (rw *RWMutex) Unlock()
*/
func implementMutex() {

	const gr = 100
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gr)

	var mu sync.Mutex

	for i := 0; i < gr; i++ {

		go func() {
			mu.Lock()
			v := counter // v is local for each goroutine
			// time.Sleep(time.Second)
			// runtime.Gosched()   /runtime.Gosched() does not guarantee that all goroutines will execute in order.
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()

		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	fmt.Println("Counter : ", counter)
}
