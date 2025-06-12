package input

type AddUserInput struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}
