package auth

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserConfirm struct {
	User             User   `json:"user"`
	ConfirmationCode string `json:"confirmation_code"`
}

type UserForgot struct {
	Email string `json:"email"`
}
