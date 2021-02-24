package model

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserDto struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type SigninRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SigninResponse struct {
	Token     string `json:"token"`
	TokenType string `json:"token_type"`
}
