package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println(".... ===> Your Profit Calculator <=== ....")

	// var revenue float64
	// var expenses float64
	// var taxRate float64

	revenue, err1 := getUserInput("Revenue: ")
	// if err1 != nil {
	// 	fmt.Println("Error found: ", err1)
	// 	return
	// }
	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	expenses, err2 := getUserInput("Expenses: ")
	// if err2 != nil {
	// 	fmt.Println("Error found: ", err2)
	// 	return
	// }
	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	taxRate, err3 := getUserInput("Tax Rate: ")
	// if err3 != nil {
	// 	fmt.Println("Error found: ", err3)
	// 	return
	// }
	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Error found: ", err1)
		return
	}

	ebt, profit, ratio := calulateValues(revenue, expenses, taxRate)
	writeCalulatedData(ebt, profit, ratio)
	// ebt := revenue - expenses
	// profit := ebt * (1 - taxRate/100)
	// ratio := ebt / profit

	fmt.Printf("Earnings Before Tax (EBT) ===> %.2f\n", ebt)

	fmt.Printf("Earnings After Tax (Profit) ===> %.2f\n", profit)

	fmt.Printf("Ratio (EBT)/Profit ===> %.2f\n", ratio)

	fmt.Print("Profit Calculated successfully !")
}

func getUserInput(text string) (float64, error) {
	var value float64
	fmt.Print(text)
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("Error, value can not be less then 0.")
	}
	return value, nil
}

func calulateValues(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func writeCalulatedData(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.2f \n Profit: %.2f \n Ratio: %.2f \n", ebt, profit, ratio)
	os.WriteFile("data.txt", []byte(results), 0664)
}
