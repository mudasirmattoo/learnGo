package main

import "fmt"

func main() {

	//value semantics
	a := 1
	fmt.Println(a)         //1
	fmt.Println(addOne(a)) //2
	fmt.Println(a)         //1

	b := 1
	fmt.Println(b) //1
	addOneP(&b)
	fmt.Println(b) //2
}

//value semantics ---> creates a completely independent copy of a variable
func addOne(v int) int {
	return v + 1
}

//pointer semantics
func addOneP(v *int) {
	*v = *v + 1
}
