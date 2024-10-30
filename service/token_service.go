package service

import (
	"online_chat/enviroment"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAccessToken(username string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
        "exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	access_token, _ := token.SignedString([]byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")))

	return access_token
}


func NewRefreshToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 168).Unix(),
		})

	refresh_token, _ := token.SignedString([]byte(enviroment.GoDotEnvVariable("REFRESH_TOKEN_SECRET")))

	return refresh_token
}

func ExtractUsernameFromToken(header string) string {
	header_parts := strings.Split(header, " ")
	token, _ := jwt.Parse(header_parts[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	var username string
	if ok && token.Valid {
        username = claims["username"].(string)	
    }
	return username
}