package entity

type User struct {
	UserID   int64  `json:"user_id" example:"1" validate:"required"`
	Username string `json:"username" example:"username"`
	Email    string `json:"email" example:"username@email.com" validate:"required"`
	//Login        string `json:"login"`
	//PasswordHash string `json:"password_hash"`
	//PasswordSalt string `json:"password_salt"`
	//Secret       string `json:"secret"`
}
