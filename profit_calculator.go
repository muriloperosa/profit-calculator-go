package main

import (
	"fmt"
)

func main() {
	revenue := input("Revenue: ")
	expenses := input("Expenses: ")
	taxRate := input("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Print("\n========\n\n")

	output("EBT: ", ebt)
	output("Profit: ", profit)
	output("Ratio: ", ratio)
}

func input(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calculateFinancials(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func output(text string, value float64) {
	fmt.Printf(text + fmt.Sprintf("%.2f\n", value))
}
