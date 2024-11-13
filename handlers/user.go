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
	"gorm.io/gorm/clause"
)

func CreateUser(c echo.Context) error {
    username, email, password := c.FormValue("username"), c.FormValue("email"), c.FormValue("password")
    
    err_message := validation.Validate(username, email, password)
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
		"status": 0,
		"tokens": map[string]string{
			"access_token": service.NewAccessToken(username),
            "refresh_token": service.NewRefreshToken(username),
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
    username := service.ExtractUsernameFromToken(token, enviroment.GoDotEnvVariable("ACCESS_TOKEN_SECRET"))

    var user models.User
    db.Preload(clause.Associations).Where("username = ?", username).Find(&user)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": 0,
        "user": user,
    })
}

func GetUserByID(c echo.Context) error {
    var user models.User
    db.Preload(clause.Associations).Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.JSON(http.StatusNotFound, map[string]string{
            "status": "1",
            "error": "no user found",
        })
    }

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": 0,
        "user": user,
    })
}

func UpdateUser(c echo.Context) error {
    var user models.User
    db.Preload(clause.Associations).Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.JSON(http.StatusNotFound, map[string]string{
            "status": "1",
            "error": "no user found",
        })
    }

    user.Username, user.Email, user.Password.Hash = c.FormValue("username"), c.FormValue("email"), password_hashing.HashPassword(c.FormValue("password"))

    db.Save(&user)

    return c.JSON(http.StatusOK, map[string]interface{}{
        "status": 0,
        "user": user,
    })
}

func DeleteUser(c echo.Context) error {
    var user models.User
    db.Preload("Password").Preload("Room").Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.JSON(http.StatusNotFound, map[string]string{
            "status": "1",
            "error": "no user found",
        })
    }

    db.Delete(&user)

    return c.JSON(http.StatusOK, map[string]string{
        "status": "0",
        "message": "user deleted",
    })
}



