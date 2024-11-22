package service

import "online_chat/models"

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

	if password_form.Hash != "" {
		fields["hash"] = password_form.Hash
	}

	return fields
}