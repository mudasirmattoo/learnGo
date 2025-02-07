package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)

	fmt.Println(createdDate)
}
