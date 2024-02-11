package res

type UserRes struct {
	ID uint    `json:"id"`
	Email string `json:"email"`
	FullName string `json:"full_name"`
	UserName string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Token string `json:"token"`
}