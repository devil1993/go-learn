package main

import (
	"fmt"
)

func main() {
	var revenue, expenses int
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rates: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := float64(ebt) * (1 - float64(taxRate)/100)
	ratio := float64(ebt) / profit

	fmt.Printf("Calculated EBT = %v\n", ebt)
	fmt.Println("Calculated Profit = ", profit)
	formattedRatio := fmt.Sprintf("Calculated Ratio = [%.2f]\n", ratio)

	fmt.Println(formattedRatio)

}
