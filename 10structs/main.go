package main

import "fmt"

func main() {

	//no inheritance in golang, no Super or parent

	mudasir := User{"mudasir", "mudasir@gmail.com", true, 17}
	fmt.Println(mudasir)                                                                //{mudasir mudasir@gmail.com true 17}
	fmt.Printf("mudasir's details are : %+v\n", mudasir)                                //mudasir's details are : {Name:mudasir Email:mudasir@gmail.com status:true age:17}
	fmt.Printf("mudasir's name is %+v and email is %+v\n", mudasir.Name, mudasir.Email) //mudasir's name is mudasir and email is mudasir@gmail.com
}

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}
