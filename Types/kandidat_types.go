package types

import models "PemilihanAPI/Model"

// CreateKandidatRequest adalah struktur untuk request create kandidat
type CreateKandidatRequest struct {
	NomorUrut int    `json:"nomor_urut" form:"nomor_urut" validate:"required"`
	NamaKetua string `json:"nama_ketua" form:"nama_ketua" validate:"required"`
	NamaWakil string `json:"nama_wakil" form:"nama_wakil" validate:"required"`
	Visi      string `json:"visi" form:"visi" validate:"required"`
	Misi      string `json:"misi" form:"misi" validate:"required"`
	Foto      string `json:"foto" form:"foto"`
}

// UpdateKandidatRequest adalah struktur untuk request update kandidat
type UpdateKandidatRequest struct {
	NomorUrut int     `json:"nomor_urut" form:"nomor_urut"`
	NamaKetua string  `json:"nama_ketua" form:"nama_ketua"`
	NamaWakil string  `json:"nama_wakil" form:"nama_wakil"`
	Visi      string  `json:"visi" form:"visi"`
	Misi      string  `json:"misi" form:"misi"`
	Foto      *string `json:"foto" form:"foto"` // Pointer untuk membedakan "tidak diisi" (nil) dan "diisi dengan kosong" (pointer ke "")
}

// KandidatResponse adalah struktur untuk response kandidat
type KandidatResponse struct {
	Message string            `json:"message"`
	Data    *models.Kandidat  `json:"data,omitempty"`
}

// KandidatListResponse adalah struktur untuk response list kandidat
type KandidatListResponse struct {
	Message string              `json:"message"`
	Data    []models.Kandidat  `json:"data"`
}
