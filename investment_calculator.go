package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting program ....")

	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var noOfYears float64 = 10

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100,noOfYears)

	fmt.Println("Future Value ===>", futureValue)

	fmt.Print("Program ended successfully !")
}