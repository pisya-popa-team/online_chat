package handlers

import (
	"net/http"
	"online_chat/email"
	"online_chat/service"
	"online_chat/utils"
	"time"

	"online_chat/models"
	"online_chat/validation"

	"github.com/labstack/echo/v4"
)

func GenerateRecoveryToken(c echo.Context) error {
	user_email := c.FormValue("user_email")
	err_message := validation.ValidateOther(user_email, "", "")
    if (err_message != "") {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "status": "1",
            "error": err_message,
        })
    }

	code := utils.IntToString(service.RecoveryToken())

	var user models.User
	db.Where("email =?", user_email).Find(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
            "status": "1",
            "error": "user not found",
        })
	}

	if email.EmailSender(user_email, code) != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
            "status": "1",
            "error": "error",
        })
	}

	var recovery models.Recovery
	db.Where("user_id =?", user.ID).Find(&recovery)

	if recovery.ID == 0 {
		recovery := models.Recovery{
			Token: code,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Hour * 1),
			UserID: user.ID,
		}
	
		db.Create(&recovery)

		return c.JSON(http.StatusOK, map[string]string{
			"status": "0",
			"message": "recovery code sent",
		})
	}

	db.Model(&recovery).Updates(
		models.Recovery{
			Token: code,
            CreatedAt: time.Now(),
            ExpiresAt: time.Now().Add(time.Hour * 1),
			IsUsed: false,
		},
	)

	return c.JSON(http.StatusOK, map[string]string{
        "status": "0",
        "message": "recovery code sent",
    })
}