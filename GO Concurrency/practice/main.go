package main

import (
	"fmt"
	"sync"
)

func Adder(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	var r = a + b
	fmt.Println("addition is :", r)
}
func Subtract(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	var r = a - b
	fmt.Println("Subtraction is :", r)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go Adder(5, 6, &wg)
	// time.Sleep(1 * time.Second)
	wg.Wait()
	wg.Add(1)
	Subtract(5, 6, &wg)
}
