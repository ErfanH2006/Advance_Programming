package services

import (
	"encoding/json"
	"errors"
	"finalExam/controllers/input"
	"finalExam/storage"

	"github.com/google/uuid"
)

const transactionsFile = "data/transactions.json"

var Incomes []input.Income
var Expenses []input.Expense

func LoadTransactions() error {
	// بارگذاری درآمدها و هزینه‌ها از فایل‌ها
	var raw []map[string]interface{}
	err := storage.LoadJSON(transactionsFile, &raw)
	if err != nil {
		return err
	}

	Incomes = []input.Income{}
	Expenses = []input.Expense{}

	for _, item := range raw {
		tType, ok := item["type"].(string)
		if !ok {
			continue
		}
		bytes, _ := json.Marshal(item)
		switch tType {
		case "income":
			var inc input.Income
			json.Unmarshal(bytes, &inc)
			Incomes = append(Incomes, inc)
		case "expense":
			var exp input.Expense
			json.Unmarshal(bytes, &exp)
			Expenses = append(Expenses, exp)
		}
	}

	return nil
}

func SaveTransactions() error {
	// ذخیره هر دو در یک فایل با نوع تراکنش مشخص
	var raw []map[string]interface{}

	for _, inc := range Incomes {
		m := map[string]interface{}{
			"type":        "income",
			"id":          inc.ID,
			"amount":      inc.Amount,
			"date":        inc.Date,
			"description": inc.Description,
			"user_id":     inc.UserID,
			"source":      inc.Source,
		}
		raw = append(raw, m)
	}

	for _, exp := range Expenses {
		m := map[string]interface{}{
			"type":        "expense",
			"id":          exp.ID,
			"amount":      exp.Amount,
			"date":        exp.Date,
			"description": exp.Description,
			"user_id":     exp.UserID,
			"category":    exp.Category,
		}
		raw = append(raw, m)
	}

	return storage.SaveJSON(transactionsFile, raw)
}

// ثبت درآمد جدید
func AddIncome(userID string, amount float64, date, desc, source string) (input.Income, error) {
	inc := input.Income{
		Transaction: input.Transaction{
			ID:          uuid.New().String(),
			Amount:      amount,
			Date:        date,
			Description: desc,
			UserID:      userID,
		},
		Source: source,
	}
	Incomes = append(Incomes, inc)
	err := SaveTransactions()
	return inc, err
}

// ثبت هزینه جدید
func AddExpense(userID string, amount float64, date, desc, category string) (input.Expense, error) {
	exp := input.Expense{
		Transaction: input.Transaction{
			ID:          uuid.New().String(),
			Amount:      amount,
			Date:        date,
			Description: desc,
			UserID:      userID,
		},
		Category: category,
	}
	Expenses = append(Expenses, exp)
	err := SaveTransactions()
	return exp, err
}

// حذف تراکنش بر اساس شناسه
func DeleteTransaction(id string) error {
	found := false
	for i, inc := range Incomes {
		if inc.ID == id {
			Incomes = append(Incomes[:i], Incomes[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		for i, exp := range Expenses {
			if exp.ID == id {
				Expenses = append(Expenses[:i], Expenses[i+1:]...)
				found = true
				break
			}
		}
	}
	if !found {
		return errors.New("تراکنش یافت نشد")
	}
	return SaveTransactions()
}

// مشاهده همه تراکنش‌ها
func ListAllTransactions() []interface{} {
	var result []interface{}
	for _, inc := range Incomes {
		result = append(result, inc)
	}
	for _, exp := range Expenses {
		result = append(result, exp)
	}
	return result
}