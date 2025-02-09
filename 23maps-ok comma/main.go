package main

import "fmt"

func main() {

	mpp := map[string]int{
		"mudasir": 23,
		"mushtaq": 47,
	}

	fmt.Println(mpp)

	mp := make(map[string]int)

	mp["hello"] = 1
	mp["world"] = 5
	fmt.Println(mp)

	for k, v := range mp {
		fmt.Println(k, v)
	}

	for _, v := range mpp {
		fmt.Println(v)
	}

	//ok checks if "mudasir" is present in the map , if ok print its value
	age, ok := mpp["mudasir"]
	if ok {
		fmt.Println(age)
	}

	age1, ok := mpp["helllllll"]
	if ok {
		fmt.Println(age1)
	}

	if v, ok := mpp["mushtaq"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("value doesn't exist")
	}
	//delete element
	delete(mp, "world")
	delete(mp, "world") //compiler won't panic

	fmt.Println(mp)

}
