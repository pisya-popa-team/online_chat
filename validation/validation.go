package validation

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Username string `validate:"min=3"`
	Email    string `validate:"email"`
 	Password string `validate:"password"`
}

func Validate(username string, email string, password string) string {
	validate := validator.New()
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`^(.{0,6}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
        return !re.MatchString(fl.Field().String())
    })

	user := User{
		Username: username,
        Email:    email,
        Password: password,
	}

	err := validate.Struct(user)

	error_message := ""

	if err != nil {
		error_message = err.Error()
    }

	return error_message
}