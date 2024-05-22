package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Println("\n========")

	fmt.Println("EBT: " + fmt.Sprintf("%f", ebt))
	fmt.Println("Profit: " + fmt.Sprintf("%f", profit))
	fmt.Println("Ratio: " + fmt.Sprintf("%f", ratio))
}
