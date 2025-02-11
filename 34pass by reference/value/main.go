package main

import "fmt"

/*
Reference data types in Go
-- pointers
--Functions --> variable storing value returned by a function references to the address of func
--maps
--slices --> references underlying array
--channels
--interfaces
*/

func main() {

	a := 42
	fmt.Println(a) //42
	change(&a)
	fmt.Println(a) //43

	si := []int{99, 2, 3, 4}
	fmt.Println(si) //[99 2 3 4]
	sliceDelta(si)
	fmt.Println(si) //[1 2 3 4]

	//maps
	m := make(map[string]int)

	m["hello"] = 1
	fmt.Println(m) //map[hello:1]
	mapDelta(m, "hello")
	fmt.Println(m) //map[hello:4]
}

func sliceDelta(is []int) {
	is[0] = 1
}

func mapDelta(md map[string]int, s string) {
	md[s] = 4
}

func change(n *int) { //*int means pointer to integer
	*n = 43
}
