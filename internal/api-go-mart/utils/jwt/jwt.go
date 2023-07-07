package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	Login string
}

const (
	lifeTime = 1 * time.Hour
)

var key = []byte("secret")

func GenerateToken(login string) (string, error) {
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(lifeTime)),
		},
		Login: login,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, &claims)
	return token.SignedString(key)
}

func (c *Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
