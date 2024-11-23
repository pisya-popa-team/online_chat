package models

type UpdateUser struct {
	Username string
	Email    string
}

type UpdatePassword struct {
	Password string
}