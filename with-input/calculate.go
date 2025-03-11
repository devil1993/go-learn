package main

import (
	"fmt"
	"math"
)

func main() {
	inflationRate := 4.5
	var capital, interestRate, years float64

	fmt.Print("Investment amount: ")
	fmt.Scan(&capital)

	fmt.Print("Rate of interest: ")
	fmt.Scan(&interestRate)

	fmt.Print("Duration: ")
	fmt.Scan(&years)

	amount := capital * math.Pow(1+interestRate/100, years)
	adjustedAmount := amount / math.Pow(1+inflationRate/100, years)

	fmt.Println("Total amount: ", amount)
	fmt.Println("Inflation Adjusted Amount: ", adjustedAmount)
}
