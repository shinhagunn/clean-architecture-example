package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var jwtSecret []byte

type Claims struct {
	UID string `json:"uid"`
	jwt.StandardClaims
}

func GenerateToken(uid string) (string, error) {
	nowTime := time.Now()
	expiresTime := nowTime.Add(3 * time.Minute)

	claims := &Claims{
		UID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	return tokenString, errors.Wrap(err, "signed token fail")
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
