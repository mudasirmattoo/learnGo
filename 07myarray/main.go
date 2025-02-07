package main

import "fmt"

func main() {

	var fruitlist [4]string //[4] you have to provide the size compulsorily

	fruitlist[0] = "Apple"
	fruitlist[1] = "tomato"
	fruitlist[3] = "peach"

	fmt.Println("fruit list is : ", fruitlist)             //fruit list is :  [Apple tomato  peach]  bigger gap between tomato and peach, indicating missing element
	fmt.Println("fruit list length is : ", len(fruitlist)) //fruit list length is :  4 . Even though there are only 3 elements, go doesn't calculate the actual length

	var vegList = [3]string{"potato", "potata", "beans"}
	fmt.Println("veg list is :", vegList)

}
