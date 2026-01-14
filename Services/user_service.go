package services

import (
	"database/sql"
	"errors"

	db "PemilihanAPI/DB"
	models "PemilihanAPI/Model"
	types "PemilihanAPI/Types"
)

// UserService interface (khusus aksi user: voting)
type UserService interface {
	Vote(req types.VoteRequest) error
	GetMyVote(userID int64) (*types.VoteData, error)
}

type userService struct {
	db *sql.DB
}

// NewUserService membuat instance baru UserService
func NewUserService() UserService {
	return &userService{
		db: db.DB,
	}
}

// Vote digunakan user untuk memilih kandidat
func (s *userService) Vote(req types.VoteRequest) error {
	// Cek apakah user sudah voting
	var exist int
	err := s.db.QueryRow(
		"SELECT COUNT(1) FROM voting WHERE id_user = ?",
		req.UserID,
	).Scan(&exist)

	if err != nil {
		return errors.New("terjadi kesalahan pada server")
	}

	if exist > 0 {
		return errors.New("user sudah melakukan voting")
	}

	// Simpan voting
	_, err = s.db.Exec(
		"INSERT INTO voting (id_user, kandidat, status) VALUES (?, ?, ?)",
		req.UserID,
		req.Kandidat,
		"voted",
	)

	if err != nil {
		return errors.New("gagal menyimpan data voting")
	}

	return nil
}

// GetMyVote mengambil data voting milik user
func (s *userService) GetMyVote(userID int64) (*types.VoteData, error) {
	var vote models.Voting

	err := s.db.QueryRow(
		`SELECT id_vote, id_user, kandidat, status, created_at
		 FROM voting WHERE id_user = ?`,
		userID,
	).Scan(
		&vote.IDVote,
		&vote.UserID,
		&vote.Kandidat,
		&vote.Status,
		&vote.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user belum melakukan voting")
		}
		return nil, errors.New("terjadi kesalahan pada server")
	}

	return &types.VoteData{
		IDVote:    vote.IDVote,
		UserID:    vote.UserID,
		Kandidat:  vote.Kandidat,
		Status:    vote.Status,
		CreatedAt: vote.CreatedAt,
	}, nil
}
