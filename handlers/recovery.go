package handlers

import (
	"net/http"
	"online_chat/email"
	"online_chat/password_hashing"
	"online_chat/service"
	"online_chat/utils"
	"time"

	"online_chat/models"
	"online_chat/validation"

	"github.com/labstack/echo/v4"
)

func GenerateRecoveryToken(c echo.Context) error {
	user_email := c.FormValue("user_email")
	err_message := validation.Validate(validation.User{
		Email: user_email,
	}, 
	validation.Options{
        Tag: utils.PointerTo("email"),
    })

    if (err_message != "") {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "status": "2",
            "error": err_message,
        })
    }

	var (
		user models.User
		recovery models.Recovery
	)

	code := service.RecoveryToken()

	if code == "" {
		return c.JSON(http.StatusInternalServerError, map[string]string{
            "status": "4",
            "error": "error with generation",
        })
	}

	db.Where("email =?", user_email).Find(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
            "status": "3",
            "error": "user not found",
        })
	}

	if email.EmailSender(user_email, code) != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
            "status": "5",
            "error": "error with sending",
        })
	}
	
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

	db.Model(&recovery).Updates(map[string]interface{}{
		"token": code,
		"created_at": time.Now(),
		"expires_at": time.Now().Add(time.Hour * 1),
		"is_used": false,
	},
	)

	return c.JSON(http.StatusOK, map[string]string{
        "status": "0",
        "message": "recovery code sent",
    })
}

func GetRecoveryToken(c echo.Context) error {
	token := c.Param("token")
	var recovery models.Recovery
	db.Where("token =?", token).Find(&recovery)
	if recovery.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
            "status": "3",
            "error": "recovery token not found",
        })
	}

	return c.JSON(http.StatusOK, map[string]string{
        "status": "0",
        "message": "recovery token found",
    })
}

func UseRecoveryToken(c echo.Context) error {
	user_token := c.FormValue("token")
	user_password := c.FormValue("password")

	var recovery models.Recovery
	db.Where("token =?", user_token).Find(&recovery)

	if recovery.ID == 0 || recovery.IsUsed || recovery.ExpiresAt.Before(time.Now()) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"status": "1",
			"message": "recovery token is invalid",
		})
	}

	err_message := validation.Validate(validation.User{
		Password: user_password,
	}, validation.Options{
        Tag: utils.PointerTo("password"),
    })

    if err_message != "" {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "status": "2",
            "error": err_message,
        })
    }

	var password models.Password
	db.Where("user_id = ?", recovery.UserID).Find(&password)

	db.Model(&password).Update("hash", password_hashing.HashPassword(user_password))
	db.Model(&recovery).Update("is_used", true)

	return c.JSON(http.StatusOK, map[string]string{
		"status": "0",
        "message": "password updated",
	})
}

