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

	// Kandidat routes
	e.POST("/kandidat", controllers.CreateKandidat)
	e.GET("/kandidat", controllers.GetAllKandidat)
	e.GET("/kandidat/:id", controllers.GetKandidatByID)
	e.PUT("/kandidat/:id", controllers.UpdateKandidat)
	e.DELETE("/kandidat/:id", controllers.DeleteKandidat)

	// Voting routes
	e.POST("/voting", controllers.Vote)           // user melakukan voting
	e.GET("/voting", controllers.GetAllVoting)    // list semua voting (admin)
	e.GET("/voting/result", controllers.GetVotingResult) // rekap hasil voting per kandidat
}

