package utils

import (
	"bufio"
	"finalExam/controllers/services"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var activeUserID string

func ShowMainMenu() {
	fmt.Println("\n=== Personal financial management ===")
	fmt.Println("1. Members management")
	fmt.Println("2. Record income")
	fmt.Println("3. Cost registration")
	fmt.Println("4. View transactions")
	fmt.Println("5. Reports")
	fmt.Println("0. Exit")
}

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func HandleUserMenu() {
	for {
		fmt.Println("\n--- Members management ---")
		fmt.Println("1. Add member")
		fmt.Println("2. View users")
		fmt.Println("3. Change member status")
		fmt.Println("0. Back")
		choice := GetUserInput("Choice: ")

		switch choice {
		case "1":
			username := GetUserInput("Username: ")
			email := GetUserInput("Email: ")
			user, err := services.AddUser(username, email)
			if err != nil {
				fmt.Println("Error in Add member proces:", err)
			} else {
				if activeUserID == "" {
					activeUserID = user.ID
					services.LoadDefaultCategories(user.ID)
				}
				fmt.Println("user has sucessfully added.")
			}
		case "2":
			users := services.ListUsers()
			fmt.Println("\nMembers List: ")
			for _, u := range users {
				activeMark := ""
				if u.ID == activeUserID {
					activeMark = "(Active)"
				}
				fmt.Printf("- %s %s - %s %s\n", u.UserName, u.Email, u.ID, activeMark)
			}
		case "3":
			id := GetUserInput("New user ID in order to get active: ")
			_, err := services.FindUserByID(id)
			if err != nil {
				fmt.Println("User not found.")
			} else {
				activeUserID = id
				fmt.Println("Active user has correctly changed.")
			}
		case "0":
			return
		default:
			fmt.Println("Invalid Option.")
		}
	}
}

func HandleAddIncome() {
	if activeUserID == "" {
		fmt.Println("You should choose an active user.")
		return
	}
	amountStr := GetUserInput("amount of income: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount.")
		return
	}
	source := GetUserInput("source of income: ")
	date := GetUserInput("Date (YYYY-MM-DD) (it will be for today if you leave it empty): ")
	if date == "" {
		date = time.Now().Format("YYYY-MM-DD")
	}
	desc := GetUserInput("Description: ")

	_, err = services.AddIncome(activeUserID, amount, date, desc, source)
	if err != nil {
		fmt.Println("Error in Recording Income :", err)
	} else {
		fmt.Println("Income has cusccessfully registered.")
	}
}

func HandleAddExpense() {
	if activeUserID == "" {
		fmt.Println("You should choose an active member first.")
		return
	}

	amountStr := GetUserInput("Amount of charge: ")
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		fmt.Println("Invalid amount.")
		return
	}

	categories := services.GetUserCategories(activeUserID)
	fmt.Println("Available categories:")
	for i, c := range categories {
		fmt.Printf("%d. %s\n", i+1, c.Name)
	}
	fmt.Println("0. Add new Category")
	choice := GetUserInput("Choose a category ")
	var categoryName string
	if choice == "0" {
		categoryName = GetUserInput("name of category: ")
		err := services.AddCategory(activeUserID, categoryName)
		if err != nil {
			fmt.Println("Error in adding new category:", err)
			return
		}
		fmt.Println("New Category added successfully.")
	} else {
		index, err := strconv.Atoi(choice)
		if err != nil || index < 1 || index > len(categories) {
			fmt.Println("invalid category.")
			return
		}
		categoryName = categories[index-1].Name
	}

	date := GetUserInput("Date (YYYY-MM-DD) (it will be for today if you leave it empty): ")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	desc := GetUserInput("Description: ")

	_, err = services.AddExpense(activeUserID, amount, date, desc, categoryName)
	if err != nil {
		fmt.Println("Error in cost registration:", err)
	} else {
		fmt.Println("The expense was successfully recorded.")
	}
}

func HandleListTransactions() {
	if activeUserID == "" {
		fmt.Println("You should choose an active member first.")
		return
	}

	fmt.Println("\n Transactions list:")
	incomes := services.Incomes
	expenses := services.Expenses

	for _, inc := range incomes {
		if inc.UserID == activeUserID {
			fmt.Printf("[Income] Amount: %.2f، Source: %s، Date: %s، Description: %s، ID: %s\n", inc.Amount, inc.Source, inc.Date, inc.Description, inc.ID)
		}
	}
	for _, exp := range expenses {
		if exp.UserID == activeUserID {
			fmt.Printf("[Income] Amount: %.2f، Source: %s، Date: %s، Description: %s، ID: %s\n",
				exp.Amount, exp.Category, exp.Date, exp.Description, exp.ID)
		}
	}

	choice := GetUserInput("Do you want to delete a transaction? (y/n): ")
	if strings.ToLower(choice) == "y" {
		id := GetUserInput("ID of transaction that you want to be deleted: ")
		err := services.DeleteTransaction(id)
		if err != nil {
			fmt.Println("Error deleting transaction:", err)
		} else {
			fmt.Println("Transaction successfully deleted.")
		}
	}
}

func HandleReports() {
	if activeUserID == "" {
		fmt.Println("You should choose an active member first.")
		return
	}

	fmt.Println("\n--- Reports ---")
	fmt.Println("1. Monthly report")
	fmt.Println("2. Report based on category")
	fmt.Println("3. Daily balance")
	fmt.Println("0. Back")

	choice := GetUserInput("Choice: ")
	switch choice {
	case "1":
		ShowMonthlySummary()
	case "2":
		ShowCategoryReport()
	case "3":
		ShowDailyBalance()
	case "0":
		return
	default:
		fmt.Println("Invalid option.")
	}
}

func ShowMonthlySummary() {
	if activeUserID == "" {
		fmt.Println("No active user.")
		return
	}
	incomes := services.Incomes
	expenses := services.Expenses

	monthSummary := make(map[string]struct {
		Income  float64
		Expense float64
	})

	for _, inc := range incomes {
		if inc.UserID != activeUserID {
			continue
		}
		month := inc.Date[:7]
		s := monthSummary[month]
		s.Income += inc.Amount
		monthSummary[month] = s
	}

	for _, exp := range expenses {
		if exp.UserID != activeUserID {
			continue
		}
		month := exp.Date[:7]
		s := monthSummary[month]
		s.Expense += exp.Amount
		monthSummary[month] = s
	}

	fmt.Println("📊 Monthly Summary:")
	for month, s := range monthSummary {
		fmt.Printf("%s | Income: %.2f | Expense: %.2f | Balance: %.2f\n", month, s.Income, s.Expense, s.Income-s.Expense)
	}
}

func ShowCategoryReport() {
	if activeUserID == "" {
		fmt.Println("No active user.")
		return
	}
	categoryTotals := make(map[string]float64)
	for _, exp := range services.Expenses {
		if exp.UserID == activeUserID {
			categoryTotals[exp.Category] += exp.Amount
		}
	}
	fmt.Println("📊 Expense by Category:")
	for category, total := range categoryTotals {
		fmt.Printf("- %s: %.2f\n", category, total)
	}
}

func ShowDailyBalance() {
	if activeUserID == "" {
		fmt.Println("No active user.")
		return
	}
	balanceByDate := make(map[string]float64)

	for _, inc := range services.Incomes {
		if inc.UserID == activeUserID {
			balanceByDate[inc.Date] += inc.Amount
		}
	}
	for _, exp := range services.Expenses {
		if exp.UserID == activeUserID {
			balanceByDate[exp.Date] -= exp.Amount
		}
	}
	fmt.Println("📅 Daily Balance:")
	for date, bal := range balanceByDate {
		fmt.Printf("- %s: %.2f\n", date, bal)
	}
}
