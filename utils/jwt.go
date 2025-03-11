package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "N4qLirZXcRUvHInb4YHoAChsgwFKUYdy"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpeced signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Could not parse token")
	}

	validToken := parsedtoken.Valid

	if !validToken {
		return errors.New("Not a valid token")
	}
	//claims, ok := parsedtoken.Claims.(jwt.MapClaims)

	//if !ok {
	//	return errors.New("Invalid token claims")
	//}

	//emaiL := claims["email"].(string)
	//userId := claims["userId"].(int64)
	return nil
}
