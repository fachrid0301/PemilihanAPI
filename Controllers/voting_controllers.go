package controllers

import (
	"log"
	"net/http"
	"strconv"

	services "PemilihanAPI/Services"
	types "PemilihanAPI/Types"

	"github.com/labstack/echo/v4"
)

var votingService services.VotingService

// InitVoting meng-inject VotingService
func InitVoting(svc services.VotingService) {
	votingService = svc
}

// Vote menangani request voting (user memilih kandidat)
func Vote(c echo.Context) error {
	var req types.CreateVotingRequest
	
	// Bind dari JSON body terlebih dahulu
	if err := c.Bind(&req); err != nil {
		// Jika bind gagal, coba ambil dari form/query
		req.IDUser = 0
		req.IDKandidat = 0
	}

	// Fallback dari form / query jika masih kosong
	if req.IDUser == 0 {
		if v := c.FormValue("id_user"); v != "" {
			if id, err := strconv.Atoi(v); err == nil {
				req.IDUser = id
			}
		}
		if req.IDUser == 0 {
			if v := c.QueryParam("id_user"); v != "" {
				if id, err := strconv.Atoi(v); err == nil {
					req.IDUser = id
				}
			}
		}
	}
	
	if req.IDKandidat == 0 {
		if v := c.FormValue("id_kandidat"); v != "" {
			if id, err := strconv.Atoi(v); err == nil {
				req.IDKandidat = id
			}
		}
		if req.IDKandidat == 0 {
			if v := c.QueryParam("id_kandidat"); v != "" {
				if id, err := strconv.Atoi(v); err == nil {
					req.IDKandidat = id
				}
			}
		}
	}

	// Validasi
	if req.IDUser == 0 {
		log.Printf("Voting request error: id_user tidak valid atau kosong. Received: %d", req.IDUser)
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "id_user wajib diisi dan harus berupa angka yang valid",
		})
	}
	if req.IDKandidat == 0 {
		log.Printf("Voting request error: id_kandidat tidak valid atau kosong. Received: %d", req.IDKandidat)
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "id_kandidat wajib diisi dan harus berupa angka yang valid",
		})
	}

	log.Printf("Voting request: id_user=%d, id_kandidat=%d", req.IDUser, req.IDKandidat)

	voting, err := votingService.Vote(req)
	if err != nil {
		log.Printf("Voting service error: %v", err)
		// Semua error dari service dianggap bad request (validasi) kecuali yg butuh dibedakan
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, types.VotingResponse{
		Message: "Voting berhasil",
		Data:    voting,
	})
}

// GetAllVoting mengembalikan seluruh data voting (untuk admin)
func GetAllVoting(c echo.Context) error {
	votings, err := votingService.GetAllVoting()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.VotingListResponse{
		Message: "Berhasil mengambil data voting",
		Data:    votings,
	})
}

// GetVotingResult mengembalikan rekap hasil voting per kandidat
func GetVotingResult(c echo.Context) error {
	results, err := votingService.GetVotingResult()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil mengambil hasil voting",
		"data":    results,
	})
}

