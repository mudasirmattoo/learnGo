package main

import (
	"fmt"
	"time"
)

func main() {
	TimeTaken(dosomething)
}

func TimeTaken(f func()) {
	start := time.Now()
	f() // wrapped around by functions
	elapsed_time := time.Since(start)
	fmt.Println(elapsed_time)
}

func dosomething() {
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
	}
}
