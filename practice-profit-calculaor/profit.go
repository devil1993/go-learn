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

	ebt, profit, ratio := calculateMetrics(float64(revenue), float64(expenses), taxRate)

	fmt.Printf("Calculated EBT = %v\n", ebt)
	fmt.Println("Calculated Profit = ", profit)
	formattedRatio := fmt.Sprintf("Calculated Ratio = [%.2f]\n", ratio)

	fmt.Println(formattedRatio)
}

func calculateMetrics(revenue, expenses, taxRate float64) (ebt, profit, ratio float64){
	ebt = revenue - expenses
	profit = float64(ebt) * (1 - float64(taxRate)/100)
	ratio = float64(ebt) / profit

	return ebt, profit, ratio
}
