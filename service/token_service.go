package service

import (
	"errors"
	"online_chat/enviroment"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewAccessToken(id string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Minute * 15).Unix(),
	})

	access_token, err := token.SignedString([]byte(enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET")))

	if err != nil {
		return ""
	}

	return access_token
}

func NewRefreshToken(id string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  id,
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

	id := claims["id"].(string)

	return id
}

func ValidateAccessToken(access_token string, secret string) error {
	if access_token == "" {
		return errors.New("token must not be empty")
	}

	token, err := ParseToken(access_token, secret)
	if err != nil || !token.Valid {
		return errors.New("token invalid")
	}

	claims, _ := token.Claims.(jwt.MapClaims)
	exp, ok := claims["exp"].(float64)
	if ok {
		if int64(exp) < time.Now().Unix() {
			return errors.New("token expired")
		}
	} else  {
		return errors.New("invalid or missing expiration time")
	}

	return nil
}