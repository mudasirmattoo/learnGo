package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to Kashmir")
	fmt.Println("Rate Kashmir between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

	// numRating = input + 1  throws error as it cannot convert 1 into string

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("you added 1 to your rating :", numRating+1)
	}
}
