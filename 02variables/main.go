package main

import "fmt"

//jwt := 2999 //can't use var-less operator outside the method or function

const LoginToken string = "abhdhsh" //since it is a public global token, L is capital. It is same as putting Public infront of const in any other lang

func main() {
	var username string = "mudasir"
	fmt.Println(username)
	fmt.Printf("variable is of type : %T \n", username)

	var isloggedin bool = false
	fmt.Println(isloggedin)
	fmt.Printf("variable is of type : %T \n", isloggedin)

	// uint takes 2^8=256 values only, uint16,uint32,uint64
	var small uint8 = 255
	fmt.Println(small)
	fmt.Printf("variable is of type : %T \n", small)

	// float32 gives just 5 decimal points
	var smallfloat float32 = 255.47267289
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type : %T \n", smallfloat)

	var smallfloat64 float64 = 255.47267289324
	fmt.Println(smallfloat64)
	fmt.Printf("variable is of type : %T \n", smallfloat64)

	// default values and some aliases

	var anothervariable int
	fmt.Println(anothervariable) //no garbage value is printed , just 0
	fmt.Printf("variable is of type : %T \n", anothervariable)

	var anotherstring string
	fmt.Println(anotherstring) //no garbage value is printed , just empty string
	fmt.Printf("variable is of type : %T \n", anotherstring)

	//implicit type declaration

	var website = "mudasir.com"
	fmt.Println(website)

	//no var style

	numberOfUsers := 39000 //Scoped, you can use var-less operator inside the method, not ouside
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type : %T \n", LoginToken)

}
