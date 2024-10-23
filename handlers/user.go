package handlers

import (
	"net/http"
	"online_chat/models"
	"online_chat/password_hashing"
	"online_chat/service"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
    
    username, email, password := c.FormValue("username"), c.FormValue("email"), c.FormValue("password")

    if email != "" && username != "" && password != "" {
        return c.String(http.StatusBadRequest, "invalid arguments")
    }

    user := models.User{
        Username: username,
        Email: email,
        Password: models.Password{
            Hash: password_hashing.HashPassword(password),
        },
    }

    access, err1 := service.NewAccessToken(username, email)
    refresh, err2 := service.NewRefreshToken()

    if err1 != nil || err2 != nil {
        return c.String(http.StatusInternalServerError, "failed to generate tokens")
    }

    db.Create(&user)

    return c.JSON(http.StatusCreated, map[string]string{
        "username": user.Username,
        "access_token": access,
        "refresh_token": refresh,
    })
}

func GetAllUsers(c echo.Context) error {
    var users []models.User
    db.Preload("Password").Find(&users)
    
    return c.JSON(http.StatusOK, users)
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



