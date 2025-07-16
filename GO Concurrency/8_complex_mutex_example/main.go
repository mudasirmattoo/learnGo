package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var m sync.Mutex

	fmt.Println(bankBalance)

	incomes := []Income{
		{Source: "Main", Amount: 400},
		{Source: "Part time", Amount: 100},
		{Source: "Investments", Amount: 200},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				m.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				m.Unlock()

				fmt.Printf("On week %d, you earned $%d from source %s\n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()
	fmt.Println(bankBalance)
}
