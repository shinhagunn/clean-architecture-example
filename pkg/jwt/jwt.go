package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/shinhagunn/todo-backend/config"
)

type Claims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

func GenerateJWTToken(uid string) (string, error) {
	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Cfg.SecretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
