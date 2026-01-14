package controllers

import (
	"net/http"
	"strconv"

	services "PemilihanAPI/Services"
	types "PemilihanAPI/Types"

	"github.com/labstack/echo/v4"
)

var userService services.UserService

// InitVote initializes user service for voting
func InitVote(svc services.UserService) {
	userService = svc
}

// Vote menangani request user melakukan voting
func Vote(c echo.Context) error {
	var req types.VoteRequest

	// Bind dari JSON / form / query
	_ = c.Bind(&req)

	// Fallback UserID
	if req.UserID == 0 {
		userIDStr := c.FormValue("user_id")
		if userIDStr == "" {
			userIDStr = c.QueryParam("user_id")
		}
		if userIDStr != "" {
			if id, err := strconv.ParseInt(userIDStr, 10, 64); err == nil {
				req.UserID = id
			}
		}
	}

	// Fallback kandidat
	if req.Kandidat == "" {
		req.Kandidat = c.FormValue("kandidat")
		if req.Kandidat == "" {
			req.Kandidat = c.QueryParam("kandidat")
		}
	}

	// Validasi
	if req.UserID == 0 || req.Kandidat == "" {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "Data tidak valid",
		})
	}

	err := userService.Vote(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, types.SuccessResponse{
		Message: "Voting berhasil",
	})
}

// GetMyVote mengambil voting milik user
func GetMyVote(c echo.Context) error {
	// Ambil user_id dari param / query / form
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		userIDStr = c.QueryParam("user_id")
	}
	if userIDStr == "" {
		userIDStr = c.FormValue("user_id")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil || userID == 0 {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "User ID tidak valid",
		})
	}

	vote, err := userService.GetMyVote(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, vote)
}
