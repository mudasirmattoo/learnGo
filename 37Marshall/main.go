// Marshalling → Converting a Go Data structures like struct to JSON/XML (serialization).
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p1 := person{"Mudasir", 23, "Male"}
	p2 := person{"Moomin", 24, "Male"}

	/*
		✔️ Converts a struct into a JSON byte slice.
		✔️ json:"name" → Customizes JSON key names.
		✔️ omitempty → Removes empty fields (Email is absent in output).
		✔️ Returns ([]byte, error)

		func Marshal(v interface{}) ([]byte, error)  //an empty interface{} is a way of defining a variable in Go
	*/

	personSlice := []person{p1, p2}
	jsonByteSlice, err := json.Marshal(personSlice) //returns a byte slice

	if err != nil {
		fmt.Println(err)
	}

	//convert the byte slice into string
	fmt.Println(string(jsonByteSlice)) //[{"Name":"Mudasir","Age":23,"Sex":"Male"},{"Name":"Moomin","Age":24,"Sex":"Male"}]
}

type person struct {
	Name string //start the field names with capital letter
	Age  int
	Sex  string
}
