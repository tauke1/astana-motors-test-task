package auth

import jwt "github.com/dgrijalva/jwt-go"

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper interface {
	GenerateToken(username string, userId uint) (signedToken string, err error)
	ValidateToken(signedToken string) (claims *JwtClaim, err error)
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Username string
	UserID   uint
	jwt.StandardClaims
}
