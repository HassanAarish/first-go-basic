package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Starting program ....")

	const inflationRate = 5.5
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	noOfYears:= 10.0

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100,noOfYears)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, noOfYears)

	fmt.Println("Future Value ===>", futureValue)

	fmt.Println("Future Real Value ===>", futureRealValue)

	fmt.Print("Program ended successfully !")
}