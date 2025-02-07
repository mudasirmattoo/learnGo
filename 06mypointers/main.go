package main

import "fmt"

func main() {

	// var ptr *int

	// fmt.Println(ptr)

	myNumber := 23

	var ptr = &myNumber //& means reference

	fmt.Println("value of actual pointer is : ", ptr)  //will print the actual memory address
	fmt.Println("value of actual pointer is : ", *ptr) // *ptr checks what is inside the pointer . It will print 23

	*ptr = *ptr * 2
	fmt.Println("new value is ", myNumber) // result is 46

	//operations are performed on the actual values, not the copies

}
