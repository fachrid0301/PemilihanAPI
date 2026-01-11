package types

// RegisterRequest adalah struktur untuk request register
type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6"`
}

// LoginRequest adalah struktur untuk request login
type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

// AuthResponse adalah struktur untuk response auth
type AuthResponse struct {
	Message string      `json:"message"`
	User    *UserData   `json:"user,omitempty"`
}

// UserData adalah struktur untuk data user di response
type UserData struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// ErrorResponse adalah struktur untuk error response
type ErrorResponse struct {
	Message string `json:"message"`
}

