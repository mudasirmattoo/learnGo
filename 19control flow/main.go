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

	//switch case

	var y int
	fmt.Println("enter y")
	fmt.Scan(&y)

	switch {
	case y < 40:
		fmt.Println("less than 40")
	case x == 40:
		fmt.Println("equal to  40")
	case x > 40:
		fmt.Println("greater than 40")
	default:
		fmt.Println("none")
	}

	var w int
	fmt.Println("enter w")
	fmt.Scan(&w)

	switch w {
	case 40:
		fmt.Println("less than 40")

	case 41:
		fmt.Println("greater than 40")
	default:
		fmt.Println("none")
	}

	switch w {
	case 40:
		fmt.Println("40")
		fallthrough
	case 41:
		fmt.Println("fallthrough 41")
		fallthrough
	case 42:
		fmt.Println("fallthrough 42")
		fallthrough
	default:
		fmt.Println("none")
	}
}
