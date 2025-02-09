package main

import "fmt"

func main() {

	//defer works in LIFO order. defer pushes the statement it is used before to the end of the function, right before the curly brackets end

	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	// Hello
	// Two
	// One
	// World

	myDefer()

}

//world,One,Two, myDefer()
//0,1,2,3,4
//hello, 43210, Two, One, World

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i) //stcak would look like bottom-->0,1,2,3,4 -->top
		//defer would unpack the stack in LIFO manner
	}
}
