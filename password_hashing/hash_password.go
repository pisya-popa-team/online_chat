package password_hashing

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string) {
	var password_byte = []byte(password)
	hashed_password, _ := bcrypt.GenerateFromPassword(password_byte, bcrypt.MinCost)
	return string(hashed_password)
}