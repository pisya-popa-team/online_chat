package handlers

import (
	"net/http"
	"online_chat/models"
	"online_chat/password_hashing"

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

    db.Create(&user)

    return c.JSON(http.StatusCreated, user)
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



