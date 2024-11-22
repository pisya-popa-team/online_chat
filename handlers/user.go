package handlers

import (
	"net/http"
	"online_chat/enviroment"
	"online_chat/models"
	"online_chat/password_hashing"
	"online_chat/service"
	"online_chat/utils"
	"online_chat/validation"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
    username, email, password := c.FormValue("username"), c.FormValue("email"), c.FormValue("password")
    
    err_message := validation.ValidateReg(username, email, password)
    if (err_message != "") {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "status": "1",
            "error": err_message,
        })
    }

    var this_user models.User
    db.Where("username = ? OR email = ?", username, email).Find(&this_user)

    if this_user.ID > 0 {
        return c.JSON(http.StatusConflict, map[string]string{
            "status": "1",
            "error": "this user already exists",
        })
    }

    user := models.User{
        Username: username,
        Email: email,
        Password: models.Password{
            Hash: password_hashing.HashPassword(password),
        },
    }

    db.Create(&user)

    return c.JSON(http.StatusCreated, map[string]interface{}{
		"status": "0",
		"tokens": map[string]string{
			"access_token": service.NewAccessToken(user.ID),
            "refresh_token": service.NewRefreshToken(user.ID),
		},
	})
}

func GetAllUsers(c echo.Context) error {
    var users []models.User
    db.Preload("Password").Preload("Room").Find(&users)
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "0",
        "users": users,
    })
}

func GetInfoAboutMe(c echo.Context) error {
    token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
    id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

    var user models.User
    db.Preload("Password").Preload("Room").Where("id = ?", id).Find(&user)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "0",
        "user": user,
    })
}

func UpdateUser(c echo.Context) error {
    token := utils.ExtractTokenFromHeaderString(c.Request().Header.Get("Authorization"))
    id := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

    var (
        user models.User
        user_password models.Password
    )
    
    db.Preload("Password").Preload("Room").Where("id = ?", id).Find(&user)
    db.Where("user_id = ?", user.ID).Find(&user_password)

    username, email, password := c.FormValue("username"), c.FormValue("email"), c.FormValue("password")
    
    err_message := validation.ValidateUpdate(username, email, password)
    if err_message != "" {
        return c.JSON(http.StatusUnprocessableEntity, map[string]string{
            "status": "1",
            "error": err_message,
        })
    }

    var other_user models.User
    db.Where("username = ? OR email = ?", username, email).Find(&other_user)

    if other_user.ID != 0 {
        return c.JSON(http.StatusConflict, map[string]string{
            "status": "1",
            "error": "this user already exists",
        })
    }

    user_form := models.UpdateUser{
        Username: username,
        Email: email,
    }

    password_form := models.UpdatePassword{
        Hash: password_hashing.HashPassword(password),
    }

    db.Model(&user).Updates(service.UpdateUserWithFields(user_form))
    db.Model(&user_password).Updates(service.UpdatePasswordWithFields(password_form))

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": "0",
        "user": user,
    })
}



