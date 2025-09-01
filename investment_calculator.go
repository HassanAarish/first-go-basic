package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting program ....")

	// Different ways to use the values an there types depending on the requirements.

	// investmentAmount, noOfYears, expectedReturnRate := 1000.00, 10.00, 5.5
	// investmentAmount:= 1000.00
	// expectedReturnRate :=10.00
	// noOfYears:= 5.5

	var investmentAmount float64 = 1000
	var expectedReturnRate  float64 =10
	var noOfYears float64 = 5.5

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100,noOfYears)

	fmt.Println("Future Value ===>", futureValue)

	fmt.Print("Program ended successfully !")
}