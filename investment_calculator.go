package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	fmt.Println("Starting program ....")

	var investmentAmount float64
	var noOfYears float64
	expectedReturnRate := 5.5

	// fmt.Print("Please enter the amount you wish to invest ==> ")
	outputText(`Please enter the amount you wish to invest ==> `)
	fmt.Scan(&investmentAmount)

	// fmt.Print("Please enter the no. of years you to wish invest for ==> ")
	outputText("Please enter the no. of years you to wish invest for ==> ")
	fmt.Scan(&noOfYears)

	// fmt.Print("Please enter the expected return rate ==> ")
	outputText("Please enter the expected return rate ==> ")
	fmt.Scan(&expectedReturnRate)

	// Returning the same values in a differet way by using return statement and creating a new function
	futureValue, futureRealValue := calulateFutureValue(investmentAmount, expectedReturnRate, noOfYears)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, noOfYears)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, noOfYears)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.2f
	// Future Real Value: %.2f`, futureValue, futureRealValue)

	// fmt.Println("Future Value ===>", futureValue)

	// fmt.Println("Future Real Value ===>", futureRealValue)

	fmt.Print("Program ended successfully !")
}

func outputText(value string) {
	fmt.Print(value)
}

func calulateFutureValue(investmentAmount, expectedReturnRate, noOfYears float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, noOfYears)
	rfv = fv / math.Pow(1+inflationRate/100, noOfYears)
	return fv, rfv
	// return
}
