/*
type Stringer interface {
    String() string  //any type with String() method with string as return type is of type Stringer
}
*/

// stringer interface implementation means that if you attach method String() string to any type, it will also be of a Stringer type
package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

//Polymorphism -->  value of type struct and count are now also of type Stringer

// attach String method to type book
func (b book) String() string {
	return fmt.Sprint("the title of book is ", b.title)
}

// create a type count
type count int

func (c count) String() string {
	return fmt.Sprint("the value is ", strconv.Itoa(int(c)))
}

// wrapper function
func Loginfo(s fmt.Stringer) {
	fmt.Println("the log 138 is ", s.String())
}

func main() {

	b := book{
		title: "12 Rules for Life",
	}

	var n count = 21

	fmt.Println(b)
	fmt.Println(n)

	//logging
	log.Println(b)
	log.Println(n)
	/*
		2025/02/10 09:46:25 the title of book is 12 Rules for Life
		2025/02/10 09:46:25 the value is 21
	*/

	Loginfo(b)
	Loginfo(n)

	/*
			the log 138 is  the title of book is 12 Rules for Life
		the log 138 is  the value is 21
	*/
}
