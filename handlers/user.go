package handlers

import (
	"net/http"
	"online_chat/models"
	"online_chat/password_hashing"
	"online_chat/service"
	"online_chat/validation"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
    username, email, password := c.FormValue("username"), c.FormValue("email"), c.FormValue("password")

    if (!validation.Validate(username, email, password)) {
        return c.String(http.StatusBadRequest, "invalid arguments")
    }

    var this_user models.User
    db.Where("username = ? OR email = ?", username, email).Find(&this_user)

    if this_user.ID > 0 {
        return c.String(http.StatusConflict, "this user already exists")
    }

    user := models.User{
        Username: username,
        Email: email,
        Password: models.Password{
            Hash: password_hashing.HashPassword(password),
        },
    }

    db.Create(&user)

    access := service.NewAccessToken(username)
    refresh := service.NewRefreshToken()

    return c.JSON(http.StatusCreated, map[string]string{
        "access_token": access,
        "refresh_token": refresh,
    })
}

func GetAllUsers(c echo.Context) error {
    var users []models.User
    db.Preload("Password").Find(&users)
    
    return c.JSON(http.StatusOK, users)
}

func GetInfoAboutMe(c echo.Context) error {
    auth_header := c.Request().Header.Get("Authorization")
    username := service.ExtractUsernameFromToken(auth_header)

    var user models.User
    db.Preload("Password").Where("username = ?", username).Find(&user)

    return c.JSON(http.StatusOK, user)
}

func GetUserByID(c echo.Context) error {
    var user models.User
    db.Preload("Password").Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.String(http.StatusNotFound, "no user found")
    }

    return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
    var user models.User
    db.Preload("Password").Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.String(http.StatusNotFound, "no user found")
    }

    user.Username, user.Email, user.Password.Hash = c.FormValue("username"), c.FormValue("email"), password_hashing.HashPassword(c.FormValue("password"))

    db.Save(&user)

    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
    var user models.User
    db.Preload("Password").Take(&user, c.Param("id"))

    if user.ID == 0 {
        return c.String(http.StatusNotFound, "no user found")
    }

    db.Delete(&user)

    return c.String(http.StatusOK, "user deleted")
}



