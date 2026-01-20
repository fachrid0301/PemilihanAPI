package services

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	db "PemilihanAPI/DB"
	models "PemilihanAPI/Model"
	types "PemilihanAPI/Types"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// AuthService adalah interface untuk auth service
type AuthService interface {
	Register(req types.RegisterRequest) error
	Login(req types.LoginRequest) (*types.UserData, error)
}

type authService struct {
	db *sql.DB
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService() AuthService {
	return &authService{
		db: db.DB,
	}
}

// Register melakukan proses registrasi user baru
func (s *authService) Register(req types.RegisterRequest) error {
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("gagal melakukan hash password")
	}

	// Cek apakah username atau email sudah ada
	var existCount int
	err = s.db.QueryRow("SELECT COUNT(1) FROM users WHERE username = ? OR email = ?", req.Username, req.Email).Scan(&existCount)
	if err != nil {
		log.Printf("Register error: cek duplikasi gagal - %v", err)
		return errors.New("terjadi kesalahan pada server: " + err.Error())
	}
	if existCount > 0 {
		return errors.New("username atau email sudah terdaftar")
	}

	// Set default role sebagai "user" jika tidak ada role yang diberikan
	role := req.Role
	// Trim whitespace jika ada
	if role != "" {
		role = strings.TrimSpace(role)
	}
	if role == "" {
		role = "user"
	}
	
	// Validasi role hanya boleh "user" atau "admin"
	if role != "user" && role != "admin" {
		return errors.New("role harus 'user' atau 'admin'")
	}
	
	// Insert user ke database dengan role
	_, err = s.db.Exec(
		"INSERT INTO users (username, email, password, role) VALUES (?, ?, ?, ?)",
		req.Username, req.Email, string(hash), role,
	)

	if err != nil {
		// detect MySQL duplicate entry error (1062)
		if me, ok := err.(*mysql.MySQLError); ok && me.Number == 1062 {
			return errors.New("username atau email sudah terdaftar")
		}
		log.Printf("Register error: insert user gagal - %v", err)
		return errors.New("terjadi kesalahan pada server: " + err.Error())
	}

	return nil
}

// Login melakukan proses login user
func (s *authService) Login(req types.LoginRequest) (*types.UserData, error) {
	var user models.User
	
	// Query user dari database
	query := "SELECT id_user, username, email, password, role, created_at FROM users WHERE username = ? OR email = ?"
	log.Printf("Login attempt: username/email=%s", req.Username)
	
	err := s.db.QueryRow(query, req.Username, req.Username).Scan(
		&user.IDUser, 
		&user.Username, 
		&user.Email, 
		&user.Password, 
		&user.Role, 
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Login failed: user tidak ditemukan untuk username/email=%s", req.Username)
			return nil, errors.New("username atau email tidak ditemukan")
		}
		// Log error detail untuk debugging
		log.Printf("Login error: query database gagal - %v | Query: %s | Username: %s", err, query, req.Username)
		
		// Cek apakah error karena kolom tidak ada
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			log.Printf("MySQL Error Code: %d, Message: %s", mysqlErr.Number, mysqlErr.Message)
			if mysqlErr.Number == 1054 { // Unknown column
				return nil, errors.New("struktur database tidak sesuai. Pastikan kolom id_user, username, email, password, role, created_at ada di tabel users")
			}
		}
		
		return nil, errors.New("terjadi kesalahan pada server: " + err.Error())
	}

	log.Printf("Login: user ditemukan - ID=%d, Username=%s, Role=%s", user.IDUser, user.Username, user.Role)

	// Cek password
	// Cek apakah password sudah di-hash (bcrypt hash selalu dimulai dengan $2a$, $2b$, atau $2y$)
	isHashed := len(user.Password) >= 4 && (
		user.Password[:4] == "$2a$" || 
		user.Password[:4] == "$2b$" || 
		user.Password[:4] == "$2y$")
	
	if isHashed {
		// Password sudah di-hash dengan bcrypt
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if err != nil {
			log.Printf("Login failed: password salah untuk user ID=%d (bcrypt)", user.IDUser)
			return nil, errors.New("password salah")
		}
	} else {
		// Password masih plain text (untuk backward compatibility atau data lama)
		log.Printf("Login warning: password user ID=%d masih plain text, sebaiknya di-hash", user.IDUser)
		if user.Password != req.Password {
			log.Printf("Login failed: password salah untuk user ID=%d (plain text)", user.IDUser)
			return nil, errors.New("password salah")
		}
		// Jika plain text match, tetap berhasil login tapi log warning
		log.Printf("Login berhasil dengan plain text password - user ID=%d (disarankan update password)", user.IDUser)
	}

	log.Printf("Login berhasil: user ID=%d, Username=%s", user.IDUser, user.Username)

	// Return user data tanpa password
	return &types.UserData{
		ID:        user.IDUser,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}
