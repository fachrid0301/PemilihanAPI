package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	// Ambil environment variables dengan default values
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "pemilihan")

	// Buat connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	log.Printf("Mencoba menghubungkan ke database: %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi database: %v", err)
	}

	// Test koneksi database
	if err = DB.Ping(); err != nil {
		log.Printf("Gagal menghubungkan ke database: %v", err)
		log.Println("\n=== TROUBLESHOOTING ===")
		log.Printf("Pastikan MySQL server berjalan di %s:%s", dbHost, dbPort)
		log.Println("Cara menjalankan MySQL:")
		log.Println("  - XAMPP: Start MySQL service di XAMPP Control Panel")
		log.Println("  - WAMP: Start MySQL service di WAMP Control Panel")
		log.Println("  - MySQL Service: net start MySQL (sebagai Administrator)")
		log.Println("  - Docker: docker run -d -p 3306:3306 mysql")
		log.Println("========================\n")
		log.Fatal("Aplikasi dihentikan karena tidak dapat terhubung ke database")
	}

	log.Println("Koneksi database berhasil!")
}

// getEnv mengambil environment variable atau return default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
