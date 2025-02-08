package main

import (
	"fmt"

	"math/rand"
)

func init() {
	fmt.Println("init function print")
}

func main() {
	x := 42

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v, which is greater than x\n", z)
	} else {
		fmt.Printf("z is %v, which is smaller\n", z)
	}

	ages := map[string]int{
		"Mudasir": 25,
		"Mushtaq": 43,
	}

	age, ok := ages["Mudasir"] // ok checks if "Mudasir" is present in the map

	if ok {
		fmt.Printf("age of Mudasir : %v\n", age)
	} else {
		fmt.Println("age not found in map")
	}

	age1, ok := ages["Charlie"]
	if ok {
		fmt.Printf("age of Charlie : %v\n", age1)
	} else {
		fmt.Println("age not found in map")
	}
}
