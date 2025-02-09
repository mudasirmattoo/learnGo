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

	name := muda("Muda")
	fmt.Println(name)

	x := variadicSum(1, 2, 3, 45, 67, 7)
	fmt.Println(x)

	//unfurling slice
	xi := []int{2, 4, 5, 67, 2, 1}
	hey, msg := proAdder(xi...)
	fmt.Println(hey, msg)

}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

//variadic functions
func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hey profunction"
}

func variadicSum(sliceofInt ...int) int {
	fmt.Println(sliceofInt)

	sum := 0

	for _, v := range sliceofInt {
		sum += v
	}
	return sum
}

func greetings() {
	fmt.Println("hello there")
}

func muda(s string) string {
	return fmt.Sprint("they call me ", s)
}
