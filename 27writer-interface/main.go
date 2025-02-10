package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s ", err)
	}

	defer f.Close()

	// byte is mutable while string is not, so create a slice of byte
	s := []byte("hello akshit")

	//write the byte into the file
	_, err = f.Write(s)
	if err != nil {
		log.Fatalln(err)
	}

	//buffer

	b := bytes.NewBufferString("hi ")
	fmt.Println(b.String())
	b.WriteString("mudasir")
	fmt.Println(b.String())
	b.Reset()
	b.Write([]byte("happy"))
	fmt.Println(b.String())

}
