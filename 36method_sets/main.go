package main

import "fmt"

func main() {
	//can be passed both as a value semantic to an iinterface method
	d1 := dog{
		name: "chin",
	}

	//can be passed both as a pointer semantic and value semantic to an iinterface method
	d2 := &dog{
		name: "Chang",
	}

	d1.walk() //chin can walk
	d1.run()  //chin can run

	d2.walk() //Chang can walk
	d2.run()  //Chang can run

	// youngRun(d1)  //cannot use d1 (variable of type dog) as young value in argument to youngRun: dog does not implement young (method run has pointer receiver)

	youngRun(d2) //Chang can run

	fmt.Println("--------------------------")

	useMethodSet()
}

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println(d.name, "can walk")
}

func (d *dog) run() {
	fmt.Println(d.name, "can run")
}

type young interface {
	walk()
	run()
}

func youngRun(y young) {
	y.run()
}
