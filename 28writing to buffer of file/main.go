package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	name string
}

// any method that implements Writer interface --> default in os package
func (p person) WriteOut(w io.Writer) {
	w.Write([]byte(p.name))
}
func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s ", err)
	}

	defer f.Close()

	p := person{
		name: "mudasir",
	}

	var b bytes.Buffer

	//write to file or buffer using the same writer interface
	p.WriteOut(f)  // write to the file
	p.WriteOut(&b) //write the the buffer

	fmt.Println(b.String())
}
