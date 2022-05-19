package auth

type User struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserConfirm struct {
	Email string `json:"email"`
	ConfirmationCode string `json:"confirmation_code"`
}