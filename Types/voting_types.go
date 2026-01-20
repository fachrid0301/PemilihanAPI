package types

import models "PemilihanAPI/Model"

// CreateVotingRequest adalah body saat user melakukan voting
type CreateVotingRequest struct {
	IDUser     int `json:"id_user" form:"id_user"`
	IDKandidat int `json:"id_kandidat" form:"id_kandidat"`
}

// VotingResponse untuk response single voting
type VotingResponse struct {
	Message string         `json:"message"`
	Data    *models.Voting `json:"data,omitempty"`
}

// VotingListResponse untuk response list voting
type VotingListResponse struct {
	Message string          `json:"message"`
	Data    []models.Voting `json:"data"`
}

// VotingResult merepresentasikan rekap jumlah suara per kandidat
type VotingResult struct {
	IDKandidat int    `json:"id_kandidat"`
	NomorUrut  int    `json:"nomor_urut"`
	NamaKetua  string `json:"nama_ketua"`
	NamaWakil  string `json:"nama_wakil"`
	JumlahSuara int   `json:"jumlah_suara"`
}

