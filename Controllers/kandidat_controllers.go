package controllers

import (
	"net/http"
	"strconv"

	services "PemilihanAPI/Services"
	types "PemilihanAPI/Types"

	"github.com/labstack/echo/v4"
)

var kandidatService services.KandidatService

// InitKandidat initializes the kandidat service
func InitKandidat(svc services.KandidatService) {
	kandidatService = svc
}

// CreateKandidat menangani request create kandidat
func CreateKandidat(c echo.Context) error {
	var req types.CreateKandidatRequest

	// Bind dari semua sumber (JSON, form, query)
	_ = c.Bind(&req)

	// Fallback ke form/query params jika field masih kosong
	if req.NomorUrut == 0 {
		if nomorUrutStr := c.FormValue("nomor_urut"); nomorUrutStr != "" {
			if nomorUrut, err := strconv.Atoi(nomorUrutStr); err == nil {
				req.NomorUrut = nomorUrut
			}
		}
	}
	if req.NamaKetua == "" {
		req.NamaKetua = c.FormValue("nama_ketua")
		if req.NamaKetua == "" {
			req.NamaKetua = c.QueryParam("nama_ketua")
		}
	}
	if req.NamaWakil == "" {
		req.NamaWakil = c.FormValue("nama_wakil")
		if req.NamaWakil == "" {
			req.NamaWakil = c.QueryParam("nama_wakil")
		}
	}
	if req.Visi == "" {
		req.Visi = c.FormValue("visi")
		if req.Visi == "" {
			req.Visi = c.QueryParam("visi")
		}
	}
	if req.Misi == "" {
		req.Misi = c.FormValue("misi")
		if req.Misi == "" {
			req.Misi = c.QueryParam("misi")
		}
	}
	if req.Foto == "" {
		req.Foto = c.FormValue("foto")
		if req.Foto == "" {
			req.Foto = c.QueryParam("foto")
		}
	}

	// Validasi
	if req.NomorUrut == 0 || req.NamaKetua == "" || req.NamaWakil == "" || req.Visi == "" || req.Misi == "" {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "Data tidak valid. Nomor urut, nama ketua, nama wakil, visi, dan misi wajib diisi",
		})
	}

	kandidat, err := kandidatService.CreateKandidat(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, types.KandidatResponse{
		Message: "Kandidat berhasil dibuat",
		Data:    kandidat,
	})
}

// GetAllKandidat menangani request get all kandidat
func GetAllKandidat(c echo.Context) error {
	kandidats, err := kandidatService.GetAllKandidat()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.KandidatListResponse{
		Message: "Berhasil mengambil data kandidat",
		Data:    kandidats,
	})
}

// GetKandidatByID menangani request get kandidat by ID
func GetKandidatByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "ID tidak valid",
		})
	}

	kandidat, err := kandidatService.GetKandidatByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.KandidatResponse{
		Message: "Berhasil mengambil data kandidat",
		Data:    kandidat,
	})
}

// UpdateKandidat menangani request update kandidat
func UpdateKandidat(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "ID tidak valid",
		})
	}

	var req types.UpdateKandidatRequest
	_ = c.Bind(&req)

	// Fallback ke form/query params jika field masih kosong
	if req.NomorUrut == 0 {
		if nomorUrutStr := c.FormValue("nomor_urut"); nomorUrutStr != "" {
			if nomorUrut, err := strconv.Atoi(nomorUrutStr); err == nil {
				req.NomorUrut = nomorUrut
			}
		}
	}
	if req.NamaKetua == "" {
		req.NamaKetua = c.FormValue("nama_ketua")
		if req.NamaKetua == "" {
			req.NamaKetua = c.QueryParam("nama_ketua")
		}
	}
	if req.NamaWakil == "" {
		req.NamaWakil = c.FormValue("nama_wakil")
		if req.NamaWakil == "" {
			req.NamaWakil = c.QueryParam("nama_wakil")
		}
	}
	if req.Visi == "" {
		req.Visi = c.FormValue("visi")
		if req.Visi == "" {
			req.Visi = c.QueryParam("visi")
		}
	}
	if req.Misi == "" {
		req.Misi = c.FormValue("misi")
		if req.Misi == "" {
			req.Misi = c.QueryParam("misi")
		}
	}
	// Foto opsional, hanya set jika ada di form/query dengan nilai (untuk JSON sudah di-bind)
	// Jika dari JSON, req.Foto sudah terisi (bisa nil, pointer ke string, atau pointer ke string kosong)
	// Untuk form/query, hanya set jika ada nilai
	if req.Foto == nil {
		if foto := c.FormValue("foto"); foto != "" {
			req.Foto = &foto
		} else if foto := c.QueryParam("foto"); foto != "" {
			req.Foto = &foto
		}
	}

	kandidat, err := kandidatService.UpdateKandidat(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.KandidatResponse{
		Message: "Kandidat berhasil diupdate",
		Data:    kandidat,
	})
}

// DeleteKandidat menangani request delete kandidat
func DeleteKandidat(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: "ID tidak valid",
		})
	}

	err = kandidatService.DeleteKandidat(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, types.ErrorResponse{
		Message: "Kandidat berhasil dihapus",
	})
}
