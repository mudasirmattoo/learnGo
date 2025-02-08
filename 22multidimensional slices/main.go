package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4}
	xj := []int{5, 6, 7, 8}

	xxi := [][]int{xi, xj}

	fmt.Println(xxi)
}
