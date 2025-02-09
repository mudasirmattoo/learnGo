package main

import "fmt"

func main() {

	//no inheritance in golang, no Super or parent

	mudasir := person{"mudasir", "mudasir@gmail.com", true, 17}
	fmt.Println(mudasir)                                                                //{mudasir mudasir@gmail.com true 17}
	fmt.Printf("mudasir's details are : %+v\n", mudasir)                                //mudasir's details are : {Name:mudasir Email:mudasir@gmail.com status:true age:17}
	fmt.Printf("mudasir's name is %+v and email is %+v\n", mudasir.Name, mudasir.Email) //mudasir's name is mudasir and email is mudasir@gmail.com
	fmt.Printf("%T is the type of mudasir", mudasir)                                    //main.person is the type of mudasir{{Mushtaq hello@gmail.com true 20} Wagay true}

	//embedded structs
	sa := secreAgent{
		person: person{
			Name:   "Mushtaq",
			Email:  "hello@gmail.com",
			status: true,
			age:    20,
		},
		Name:            "Wagay",
		license_to_kill: true,
	}
	fmt.Println(sa)
	fmt.Printf("sa's details are : %+v\n", sa)
	fmt.Printf("sa's name is %+v \n", sa.Name)        //sa's name is Wagay
	fmt.Printf("sa's name is %+v \n", sa.person.Name) //sa's name is Mushtaq

	//anonymous structs
	p1 := struct {
		Name   string
		Email  string
		status bool
		age    int
	}{
		Name:   "Mushtaq",
		Email:  "hello@gmail.com",
		status: true,
		age:    20,
	}
	fmt.Printf("%T is the type of p1", p1) //struct { Name string; Email string; status bool; age int } is the type of p1

}

type person struct {
	Name   string
	Email  string
	status bool
	age    int
}

type secreAgent struct {
	person
	Name            string
	license_to_kill bool
}
