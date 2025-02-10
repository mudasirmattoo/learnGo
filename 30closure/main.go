package main

import "fmt"

func main() {

	//closure
	/*
		the incrementor function returns a function that has access to x.
		x is initialized only once per function call to incrementor().
		Every call to f() keeps modifying and retaining the state of x.
		When a new function instance is created (g := incrementor()), a new x is created, and it starts from 0 again.
	*/
	f := incrementor()
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //4
	fmt.Println(f()) //5
	fmt.Println(f()) //6

	//assigned to new function instance g,  so it will start from new memory location with 0 as initial value
	g := incrementor()
	fmt.Println(g()) //1
	fmt.Println(g()) //2
	fmt.Println(g()) //3
	fmt.Println(g()) //4
	fmt.Println(g()) //5
	fmt.Println(g()) //6
}

func incrementor() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
