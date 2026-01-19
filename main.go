package main

import (
	"log"
	"os"

	controllers "PemilihanAPI/Controllers"
	db "PemilihanAPI/DB"
	routes "PemilihanAPI/Routes"
	services "PemilihanAPI/Services"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Printf("File .env tidak ditemukan: %v, menggunakan environment variables sistem", err)
	} else {
		log.Println("File .env berhasil dimuat")
	}

	db.Connect()

	// Initialize services after database connection
	authService := services.NewAuthService()
	controllers.Init(authService)

	kandidatService := services.NewKandidatService()
	controllers.InitKandidat(kandidatService)

	e := echo.New()

	// Middleware untuk logging request dan response
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${method} ${uri} | ${latency_human} | ${error}\n",
	}))

	// Middleware untuk recover dari panic
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
    AllowOrigins: []string{"http://localhost:5173"},
    AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
    AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
}))

	// Setup semua routes
	routes.SetupRoutes(e)

	// Ambil port dari environment variable, default ke 8080
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server berjalan di port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
