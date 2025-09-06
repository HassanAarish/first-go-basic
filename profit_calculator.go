package main

import (
	"fmt"
)

func main() {
	fmt.Println(".... ===> Your Profit Calculator <=== ....")

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue := getUserInput("Revenue: ")
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses := getUserInput("Expenses: ")
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate := getUserInput("Tax Rate: ")
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calulateValues(revenue, expenses, taxRate)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("Earnings Before Tax (EBT) ===> %.2f\n", ebt)

	fmt.Printf("Earnings After Tax (Profit) ===> %.2f\n", profit)

	fmt.Printf("Ratio (EBT)/Profit ===> %.2f\n", ratio)

	fmt.Print("Profit Calculated successfully !")
}

func getUserInput(text string) float64 {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	return value
}

func calulateValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}
