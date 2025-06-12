package input

type Transaction struct {
	ID          string  `json:"id"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	UserID      string  `json:"user_id"`
}

type Income struct {
	Transaction
	Source string `json:"source"`
}

type Expense struct {
	Transaction
	Category string `json:"category"`
}

type Category struct {
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}
