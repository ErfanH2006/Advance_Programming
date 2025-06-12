package main

import (
	"finalExam/utils"
	"fmt"
)

func main() {
	for {
		utils.ShowMainMenu()
		choice := utils.GetUserInput("Choice is yours: ")
		switch choice {
		case "1":
			utils.HandleUserMenu()
		case "2":
			utils.HandleAddIncome()
		case "3":
			utils.HandleAddExpense()
		case "4":
			utils.HandleListTransactions()
		case "5":
			utils.HandleReports()
		case "0":
			fmt.Println("exit ")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
