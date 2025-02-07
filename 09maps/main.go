package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["Py"] = "Python"
	languages["Rb"] = "Ruby"
	languages["C#"] = "Csharp"

	fmt.Println("list of all the languages ", languages) //not comma separated   map[C#:Csharp JS:Javascript Py:Python Rb:Ruby]
	fmt.Println("Short of JS :", languages["JS"])        // Javascript

	delete(languages, "Rb")
	fmt.Println("list of all the languages ", languages) // map[C#:Csharp JS:Javascript Py:Python]

	//loops

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v ", key, value) //For key JS, value is Javascript For key Py, value is Python For key C#, value is Csharp
	}

	for _, value := range languages {
		fmt.Printf("value is %v ", value) //value is Csharp value is Javascript value is Python
	}
}
