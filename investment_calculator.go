package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting program ....")

	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	noOfYears := 10.0

	fmt.Print("Please enter the amount you wish to invest ==> ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter the expected return rate ==> ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Please enter the no. of years you to wish invest for ==> ")
	fmt.Scan(&noOfYears)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, noOfYears)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, noOfYears)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)

	// fmt.Printf("Future Value: %.2f\nFuture Real Value: %.2f\n", futureValue, futureRealValue)

	// fmt.Println("Future Value ===>", futureValue)

	// fmt.Println("Future Real Value ===>", futureRealValue)

	fmt.Print("Program ended successfully !")
}
