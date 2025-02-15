// package main

// func main() {
// 	f := adder(1, 2, 3, 5, 6, 6)
// 	fmt.Println(f)
// }

// func adder(v ...int) int {
// 	total := 0

// 	for _, i := range v {
// 		total += i
// 	}
// 	return total
// }

/* Interface

func main() {
	c := circle{
		radius: 4,
	}

	s := square{
		length: 34,
		width:  34,
	}

	fmt.Println(info(c))
	fmt.Println(info(s))

}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) calArea() float64 {
	return s.length * s.width
}

func (c circle) calArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	calArea() float64
}

func info(sh shape) float64 {
	return sh.calArea()
}
*/

//testing

// package main

// import (
// 	"fmt"
// 	"math"
// )

// type circle struct {
// 	radius float64
// }

// type shape interface {
// 	area() float64
// }

// func (c *circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func info(s shape) {
// 	fmt.Println("area", s.area())
// }

// /*
// The method set of a type determines the INTERFACES that the type implements.....” ~
// Receivers   Values
// -----------------------------------------------
// (t T)       T and *T   // To non pointer receiver of type T, we can pass both value and pointer of type T
// (t *T)      *T         //only poniter
// */
// func main() {
// 	c := circle{5}
// 	// info(c)      //cannot use c (variable of type circle) as shape value in argument to info: circle does not implement shape (method area has pointer receiver)
// 	fmt.Println(c.area())
// }

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// var wg sync.WaitGroup

// func main() {

// 	wg.Add(1)
// 	go hello()
// 	world()
// 	wg.Wait()
// }

//	func hello() {
//		fmt.Println("hello")
//		wg.Done()
//	}
//
//	func world() {
//		fmt.Println("world")
//	}

func main() {

	const gr = 10 //set goroutines to 10
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gr)

	for i := 0; i < gr; i++ {

		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Counter : ", counter)
			wg.Done()
		}()

		fmt.Println("Goroutines :", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines :", runtime.NumGoroutine())

	fmt.Println("Counter end value: ", counter)
}

/*Final output
different values in between
---------
Goroutines : 2
Goroutines : 3
Goroutines : 4
Goroutines : 5
Goroutines : 6
Goroutines : 7
Goroutines : 8
Goroutines : 9
Goroutines : 10
Goroutines : 11
Counter :  5
Counter :  6
Counter :  2
Counter :  7
Counter :  1
Counter :  8
Counter :  3
Counter :  9
Counter :  4
Counter :  10
Goroutines : 1
Counter end value:  10
*/
