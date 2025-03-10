package main

import (
	"fmt"
	"math"
)

func main() {
	var amount = 1000
	var expectedReturn = 5.5
	var years = 10

	var futureValue = float64(amount)*math.Pow((1+float64(expectedReturn)/100), float64(years))

	fmt.Print("Future Value: ", futureValue, "\n")
}
