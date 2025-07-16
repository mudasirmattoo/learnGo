package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func UpdateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}
func main() {

	msg = "Hello"

	var mutex sync.Mutex
	wg.Add(2)
	go UpdateMessage("Hello Mudasir!", &mutex)
	go UpdateMessage("Hello Moomin!", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
