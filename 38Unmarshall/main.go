package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	// JSON input (as a byte slice)
	// jsonData := `[{"Name":"Mudasir","Age":23,"Sex":"Male"},{"Name":"Moomin","Age":24,"Sex":"Male"}]`
	//or
	jsonData := `[{"Name":"Mudasir","Age":23,"Sex":"Male"},{"Name":"Moomin","Age":24,"Sex":"Male"}]`

	byteofslice := []byte(jsonData)

	var p []person

	//func Unmarshal(data []byte, v any) error

	err := json.Unmarshal(byteofslice, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p) //[{Mudasir 23 Male} {Moomin 24 Male}]

	/*
			Important Points About json.Unmarshal
		✔️ Converts JSON (byte slice) into a Go struct.
		✔️ Takes a pointer to the struct (&user).
		✔️ Fields must match the JSON keys (case-sensitive unless using struct tags).
		✔️ Returns error if JSON format is invalid.
	*/
}

// use json-to-go website
type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
