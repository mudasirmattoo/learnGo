package main

import (
	"fmt"
	"math"
)

// Circle struct
type circle struct {
	radius float64
}

// Interface requiring area()
type shape interface {
	area() float64
}

// Method with pointer receiver (only works on pointer)
func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Function that accepts shape interface
func info(s shape) {
	fmt.Println("Area:", s.area())
}

/*
The method set of a type determines the INTERFACES that the type implements.....
Receivers     |  Values
------------- |  ---------------------------------------
(t T)        |  T and *T   // Value receiver accepts both value & pointer
(t *T)       |  *T         // Pointer receiver accepts only pointers
*/
func useMethodSet() {
	c := circle{5}

	info(&c) // Works because area() has a pointer receiver
	//info(c)   //cannot use c (variable of type circle) as shape value in argument to info: circle does not implement shape (method area has pointer receiver)
	fmt.Println(c.area()) // This works because Go implicitly converts c to &c
}
