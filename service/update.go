package service

import (
	"online_chat/models"
	"online_chat/password_hashing"
)

func UpdateUserWithFields(user_form models.UpdateUser) interface{} {
 	fields := map[string]interface{}{
 	}

	if user_form.Username != "" {
		fields["username"] = user_form.Username
	}

	if user_form.Email != "" {
		fields["email"] = user_form.Email
	}

	return fields
}

func UpdatePasswordWithFields(password_form models.UpdatePassword) interface{} {
	fields := map[string] interface{}{	
	}

	if password_form.Password != "" {
		fields["hash"] = password_hashing.HashPassword(password_form.Password)
	}

	return fields
}