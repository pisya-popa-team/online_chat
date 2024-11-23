package validation

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"
)

type Options struct {
	Tag *string
}

type User struct {
	Username string `validate:"min=3" update:"omitempty,min=3"`
	Email    string `validate:"email" update:"omitempty,email"`
 	Password string `validate:"password" update:"omitempty,password"`
}

func InitPasswordValidation(validate *validator.Validate) {
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`^(.{0,6}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
        return !re.MatchString(fl.Field().String())
    })
}

func Validate(username string, email string, password string, options Options) string {
	validate := validator.New()
	InitPasswordValidation(validate)

	if options.Tag != nil {
		validate.SetTagName(*options.Tag)
    }

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



