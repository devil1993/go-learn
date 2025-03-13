package main

import "fmt"

func main() {
	accountBalance := 1000.0

	welcomeMessage()
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your account balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter deposit amount: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		accountBalance += depositAmount

		fmt.Println("Money deposited, your current balance is: ", accountBalance)
	} else if choice == 3 {
		fmt.Print("Enter witdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)

		accountBalance -= withdrawAmount

		fmt.Println("Money withdrawn, your current balance is: ", accountBalance)
	} else if choice == 4 {
		fmt.Println("Goodbye!!")
		return
	}

}

func welcomeMessage() {
	fmt.Println("Welcome to GO Bank")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdray money")
	fmt.Println("4. Exit")
}
