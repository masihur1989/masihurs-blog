package common

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// Claims godoc
type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// JWTSigningKey godoc
var JWTSigningKey = []byte(os.Getenv("JWT_SECRET"))

// GeneratToken godoc
func GeneratToken(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JWTSigningKey)
	if err != nil {
		l.Error(err.Error())
		return "", ErrorJWTStringGeneration
	}
	return tokenString, nil
}
