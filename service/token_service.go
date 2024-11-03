package service

import (
	"online_chat/enviroment"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAccessToken(username string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
        "exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	access_token, err := token.SignedString([]byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")))

	if err != nil {
        return ""
    }

	return access_token
}


func NewRefreshToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 168).Unix(),
		})

	refresh_token, err := token.SignedString([]byte(enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")))

	if err != nil {
        return ""
    }

	return refresh_token
}

func ParseToken(token string, secret string) (*jwt.Token, error) {
	parsed_token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return parsed_token, err
}

func ExtractUsernameFromToken(token_string string, secret string) string {
	token, _ := ParseToken(token_string, secret)

	claims, _ := token.Claims.(jwt.MapClaims)

    username := claims["username"].(string)
		
	return username
}