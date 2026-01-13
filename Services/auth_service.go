package services

import (
	"database/sql"
	"errors"
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
		return errors.New("terjadi kesalahan pada server")
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
		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}

// Login melakukan proses login user
func (s *authService) Login(req types.LoginRequest) (*types.UserData, error) {
	var user models.User
	err := s.db.QueryRow(
		"SELECT id, username, email, password, role, created_at FROM users WHERE username = ? OR email = ?",
		req.Username, req.Username,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("username atau email tidak ditemukan")
		}
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Cek password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	// Return user data tanpa password
	return &types.UserData{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}, nil
}
