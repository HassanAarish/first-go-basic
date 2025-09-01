package main

import (
	"fmt"
)

func main() {
	fmt.Println(".... ===> Your Profit Calculator <=== ....")

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
	ratio := ebt/profit

	fmt.Println("Earnings Before Tax (EBT) ===>", ebt)

	fmt.Println("Earnings After Tax (Profit) ===>", profit)

	fmt.Println("Ratio (EBT)/Profit ===>", ratio)

	fmt.Print("Profit Calculated successfully !")
}