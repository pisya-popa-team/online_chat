package service

import (
	"online_chat/enviroment"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAccessToken(username string, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
	jwt.MapClaims{
		"username": username,
		"email": email,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	access_token, err := token.SignedString([]byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return "", err
	}

	return access_token, nil
}

func NewRefreshToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 168).Unix(),
		})

	refresh_token, err := token.SignedString([]byte(enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")))

	if err != nil {
		return "", err
	}

	return refresh_token, nil
}