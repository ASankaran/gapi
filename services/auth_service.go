package services

import (
	"time"
	"fmt"
	"../models"
	"../errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService struct {
}

var secret []byte

func init() {
	secret = []byte("signing_secret")
}

func (service AuthService) GenerateToken(user *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	signed_token, _ := token.SignedString(secret)

	return signed_token
}

func (service AuthService) VerifyToken(token_string string) bool {
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int64(claims["exp"].(float64)) > time.Now().Unix() {
			return true;
		}
	}

	return false
}

func (service AuthService) GetUserID(token_string string) int {
	if !service.VerifyToken(token_string) {
		panic(errors.UnprocessableError("Bad auth token"))
	}

	token, _ := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	return int(claims["id"].(float64))
}
