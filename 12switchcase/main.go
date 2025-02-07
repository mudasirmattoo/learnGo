package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("you can open and move one spot")

	case 2:
		fmt.Println("you can move 2 spots ")

	case 3:
		fmt.Println("you can move 3 spots ")

	case 4:
		fmt.Println("you can move 4 spots ")
		fallthrough
	case 5:
		fmt.Println("you can move 5 spots ")

	case 6:
		fmt.Println("you can move 6 spots and roll dice again ")

	default:
		fmt.Println("what was that")

	}
}
