package service

import (
	"test/database/model"
	"test/interface/auth"
)

type UserService interface {
	Authenticate(username, password string) (*AuthorizationResponse, error)
	ValidateToken(token string) (*auth.JwtClaim, error)
	Get(userID uint) (*model.User, error)
	Register(username, password string) (*model.User, error)
	Logout(token string) error
}

type AuthorizationResponse struct {
	Token     string
	TokenType string
}
