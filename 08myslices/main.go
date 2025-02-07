package main

import (
	"fmt"
	"sort"
)

func main() {

	// var fruitlist = []string{}                           // in case of this syntax you need to initialize it as well by {}
	var fruitlist = []string{"apple", "peach", "lasan"}
	fmt.Printf("type of fruitless is %T\n :", fruitlist) //type of fruitless is []string

	fruitlist = append(fruitlist, "mango", "banana") // adding data to slices
	fmt.Println(fruitlist)                           //[apple peach lasan mango banana]

	fruitlist = append(fruitlist[1:])  //slicing
	fmt.Println(fruitlist)             //[peach lasan mango banana]
	fruitlist = append(fruitlist[1:3]) //start from 1 and stop at 3, 3 is not included, but 1 is included
	fmt.Println(fruitlist)             //[lasan mango]

	fruitlist = append(fruitlist[:3]) // start at defaul 0 index and stop by 3

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 //panic: runtime error: index out of range [4] with length 4

	highScores = append(highScores, 555, 654, 331) // append accomodates all the values [234 945 465 867 555 654 331]

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores) //[234 331 465 555 654 867 945]

	fmt.Println(sort.IntsAreSorted(highScores)) //true

	// how to remove a value from slices based on index

	var courses = []string{"react", "javascript", "django", "swift", "python"}
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[react javascript swift python]

}
