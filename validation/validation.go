package validation

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Username string `validate:"required"`
	Email    string `validate:"email"`
 	Password string `validate:"password"`
}



func Validate(username string, email string, password string) bool {
	validate := validator.New()
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`^(.{0,7}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
        return !re.MatchString(fl.Field().String())
    })

	user := User{
		Username: username,
        Email:    email,
        Password: password,
	}

	err := validate.Struct(user)

	return err == nil
}