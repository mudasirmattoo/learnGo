package main

import "fmt"

func main() {
	// greetings //means passing its reference
	greetings()

	func() { //anonymous function
		fmt.Println("another greetings")
	}()

	greet := func(name string) { //anonymous function
		fmt.Println("hello", name)
	}

	greet("mudasir")

	result := adder(3, 6)
	fmt.Println("result is :", result)

	proadder, mymessage := proAdder(2, 9, 20, 1, 4, 5, 7)
	fmt.Println("pro result is: ", proadder, mymessage)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hey profunction"
}

func greetings() {
	fmt.Println("hello there")
}
