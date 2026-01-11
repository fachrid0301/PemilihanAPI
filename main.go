package main

import (
	controllers "PemilihanAPI/Controllers"
	db "PemilihanAPI/DB"
	routes "PemilihanAPI/Routes"
	services "PemilihanAPI/Services"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Connect()

	// Initialize services after database connection
	authService := services.NewAuthService()
	controllers.Init(authService)

	e := echo.New()

	// Setup semua routes
	routes.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
