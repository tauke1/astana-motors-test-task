package auth

import (
	"errors"
	"test/interface/auth"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JwtWrapper wraps the signing key and the issuer
type jwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// GenerateToken generates a jwt token
func (j *jwtWrapper) GenerateToken(username string, userID uint) (signedToken string, err error) {
	claims := &auth.JwtClaim{
		Username: username,
		UserID:   userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

//ValidateToken validates the jwt token
func (j *jwtWrapper) ValidateToken(signedToken string) (claims *auth.JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&auth.JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*auth.JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return
}

func NewJwtWrapper(secretKey, issuer string, expirationHours int64) *jwtWrapper {
	if secretKey == "" {
		panic("secretKey must not be empty")
	}

	if issuer == "" {
		panic("secretKey must not be empty")
	}

	if expirationHours <= 0 {
		panic("expirationHours must be positive integer")
	}

	return &jwtWrapper{
		SecretKey:       secretKey,
		Issuer:          issuer,
		ExpirationHours: expirationHours,
	}
}
