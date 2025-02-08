package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)

	//delete 4
	xi = append(xi[:4], xi[5:]...) //... unpacks individual elements in xi[5:]
	fmt.Println(xi)
}
