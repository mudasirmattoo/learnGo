package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "Welcome to Kashmir"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating for your experience in Kashmir")

	//comma ok or comma error

	input, _ := reader.ReadString('\n') //we can use input, err for try catch or at times _, err
	fmt.Println("Thanks for rating :", input)
	fmt.Printf("The type of rating is %T", input)

}
