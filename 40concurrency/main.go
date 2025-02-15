package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("Architecture\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	// hello()
	// world()

	// hello  0
	// hello  1
	// hello  2
	// hello  3
	// hello  4
	// world  0
	// world  1
	// world  2
	// world  3
	// world  4

	// go hello()
	// world()

	// world  0
	// world  1
	// world  2
	// world  3
	// world  4

	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("using waitGroup")
	useWaitGroup()

	fmt.Println("--------atomic--------")
	implementAtomic()
	// 	OS               windows
	// Architecture             amd64
	// CPUs             4
	// Goroutines               1
	// -----------------------------------------------------------------
	// using waitGroup
	// world  0
	// world  1
	// world  2
	// world  3
	// world  4
	// CPUs     4
	// Goroutines       2
	// hello  0
	// hello  1
	// hello  2
	// hello  3
	// hello  4

	raceCondition()

	fmt.Println("-------------mutex---------------")
	implementMutex()
}

func hello() {
	for i := 0; i < 5; i++ {
		fmt.Println("hello ", i)
	}
}

func world() {
	for i := 0; i < 5; i++ {
		fmt.Println("world ", i)
	}
}
