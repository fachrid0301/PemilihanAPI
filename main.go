package main

import (
	controllers "PemilihanAPI/Controllers"
	db "PemilihanAPI/DB"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Connect()

	e := echo.New()

	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// Kalau masih mau endpoint profile tanpa proteksi, bisa:
	// e.GET("/profile", controllers.Profile)

	e.Logger.Fatal(e.Start(":8080"))
}
