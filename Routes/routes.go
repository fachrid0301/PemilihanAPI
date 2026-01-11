package routes

import (
	controllers "PemilihanAPI/Controllers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes mengatur semua routes untuk aplikasi
func SetupRoutes(e *echo.Echo) {
	// Auth routes
	e.POST("/register", controllers.Register)
	e.POST("/login", controllers.Login)

	// User routes
	// e.GET("/profile", controllers.Profile) // Jika diperlukan
}

