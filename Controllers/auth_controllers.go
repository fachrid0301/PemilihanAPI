package controllers

import (
	"net/http"

	"PemilihanAPI/DB"
	"PemilihanAPI/Model"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := db.DB.Exec(
		"INSERT INTO users (username, email, password) VALUES (?, ?, ?)",
		username, email, string(hash),
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Username atau Email sudah terdaftar",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Register berhasil",
	})
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	var user models.User
	err := db.DB.QueryRow(
		"SELECT id, username, email, password, created_at FROM users WHERE username = ?",
		username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Username atau Email tidak ditemukan",
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Password salah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login berhasil",
		"user": map[string]interface{}{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"created_at": user.CreatedAt,
		},
	})

}
