package controllers

import (
	"net/http"

	services "PemilihanAPI/Services"
	types "PemilihanAPI/Types"

	"github.com/labstack/echo/v4"
)

var authService services.AuthService

// Init initializes the auth service
func Init(svc services.AuthService) {
	authService = svc
}

// Register menangani request register
func Register(c echo.Context) error {
	var req types.RegisterRequest
	
	// Bind dari semua sumber (JSON, form, query)
	_ = c.Bind(&req)
	
	// Fallback ke form/query params jika field masih kosong
	if req.Username == "" {
		req.Username = c.FormValue("username")
		if req.Username == "" {
			req.Username = c.QueryParam("username")
		}
	}
	if req.Email == "" {
		req.Email = c.FormValue("email")
		if req.Email == "" {
			req.Email = c.QueryParam("email")
		}
	}
	if req.Password == "" {
		req.Password = c.FormValue("password")
		if req.Password == "" {
			req.Password = c.QueryParam("password")
		}
	}
	
	// Untuk role, ambil dari semua sumber dan prioritaskan yang ada
	// Echo's Bind() seharusnya sudah mengambil dari JSON, tapi pastikan dengan fallback
	roleFromForm := c.FormValue("role")
	roleFromQuery := c.QueryParam("role")
	
	// Jika role dari bind kosong, gunakan dari form/query
	if req.Role == "" {
		if roleFromForm != "" {
			req.Role = roleFromForm
		} else if roleFromQuery != "" {
			req.Role = roleFromQuery
		}
	}
	// Jika role masih kosong setelah semua, akan jadi default "user" di service
	// basic validation
	if req.Username == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "Data tidak valid",
		})
	}

	err := authService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, types.AuthResponse{
		Message: "Register berhasil",
	})
}

// Login menangani request login
func Login(c echo.Context) error {
	var req types.LoginRequest
	_ = c.Bind(&req)
	if req.Username == "" || req.Password == "" {
		if req.Username == "" {
			req.Username = c.FormValue("username")
			if req.Username == "" {
				req.Username = c.QueryParam("username")
			}
		}
		if req.Password == "" {
			req.Password = c.FormValue("password")
			if req.Password == "" {
				req.Password = c.QueryParam("password")
			}
		}
	}
	if req.Username == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "Data tidak valid",
		})
	}

	userData, err := authService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.AuthResponse{
		Message: "Login berhasil",
		User:    userData,
	})
}
