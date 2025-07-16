package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func UpdateMessage(s string) {
	defer wg.Done()
	msg = s
}
func main() {

	msg = "Hello"
	wg.Add(2)
	go UpdateMessage("Hello Mudasir!")
	go UpdateMessage("Hello Moomin!")
	wg.Wait()

	fmt.Println(msg)
}
