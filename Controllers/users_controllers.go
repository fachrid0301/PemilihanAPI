package controllers

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Profile(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"username": claims["username"],
	})
}
