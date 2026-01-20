package services

import (
	"database/sql"
	"errors"
	"log"

	db "PemilihanAPI/DB"
	models "PemilihanAPI/Model"
	types "PemilihanAPI/Types"
)

// VotingService adalah interface untuk operasi voting
type VotingService interface {
	Vote(req types.CreateVotingRequest) (*models.Voting, error)
	GetAllVoting() ([]models.Voting, error)
	GetVotingResult() ([]types.VotingResult, error)
}

type votingService struct {
	db *sql.DB
}

// NewVotingService membuat instance baru VotingService
func NewVotingService() VotingService {
	return &votingService{
		db: db.DB,
	}
}

// Vote melakukan proses voting oleh user ke kandidat tertentu
func (s *votingService) Vote(req types.CreateVotingRequest) (*models.Voting, error) {
	// Validasi user: harus ada, status aktif, dan role = 'user'
	var status string
	var role string
	err := s.db.QueryRow("SELECT status, role FROM users WHERE id_user = ?", req.IDUser).Scan(&status, &role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user tidak ditemukan")
		}
		log.Printf("Error cek user voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Hanya user dengan role 'user' yang boleh voting (admin tidak boleh)
	if role != "user" {
		return nil, errors.New("hanya user yang dapat melakukan voting")
	}

	if status == "sudah_voting" {
		return nil, errors.New("user sudah melakukan voting")
	}
	if status == "nonaktif" {
		return nil, errors.New("user tidak aktif")
	}

	// Pastikan kandidat ada
	var kandidatCount int
	err = s.db.QueryRow("SELECT COUNT(1) FROM kandidat WHERE id_kandidat = ?", req.IDKandidat).Scan(&kandidatCount)
	if err != nil {
		log.Printf("Error cek kandidat voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	if kandidatCount == 0 {
		return nil, errors.New("kandidat tidak ditemukan")
	}

	// Opsional: pastikan tidak ada record voting sebelumnya untuk user ini (backup selain status)
	var voteCount int
	err = s.db.QueryRow("SELECT COUNT(1) FROM voting WHERE id_user = ?", req.IDUser).Scan(&voteCount)
	if err != nil {
		log.Printf("Error cek voting sebelumnya: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	if voteCount > 0 {
		return nil, errors.New("user sudah melakukan voting")
	}

	// Gunakan transaction agar insert voting + update status user konsisten
	tx, err := s.db.Begin()
	if err != nil {
		log.Printf("Error mulai transaksi voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Insert ke tabel voting (waktu_voting otomatis current_timestamp)
	result, err := tx.Exec(
		"INSERT INTO voting (id_user, id_kandidat) VALUES (?, ?)",
		req.IDUser, req.IDKandidat,
	)
	if err != nil {
		tx.Rollback()
		log.Printf("Error insert voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Update status user menjadi sudah_voting
	_, err = tx.Exec("UPDATE users SET status = 'sudah_voting' WHERE id_user = ?", req.IDUser)
	if err != nil {
		tx.Rollback()
		log.Printf("Error update status user setelah voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	if err = tx.Commit(); err != nil {
		log.Printf("Error commit transaksi voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Ambil data voting yang baru dibuat
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error ambil last insert id voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	var voting models.Voting
	err = s.db.QueryRow(
		"SELECT id_voting, id_user, id_kandidat, waktu_voting FROM voting WHERE id_voting = ?",
		id,
	).Scan(&voting.IDVoting, &voting.IDUser, &voting.IDKandidat, &voting.WaktuVoting)
	if err != nil {
		log.Printf("Error ambil data voting baru: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	return &voting, nil
}

// GetAllVoting mengambil semua data voting
func (s *votingService) GetAllVoting() ([]models.Voting, error) {
	rows, err := s.db.Query("SELECT id_voting, id_user, id_kandidat, waktu_voting FROM voting ORDER BY waktu_voting DESC")
	if err != nil {
		log.Printf("Error query semua voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	defer rows.Close()

	var votings []models.Voting
	for rows.Next() {
		var v models.Voting
		if err := rows.Scan(&v.IDVoting, &v.IDUser, &v.IDKandidat, &v.WaktuVoting); err != nil {
			log.Printf("Error scan row voting: %v", err)
			return nil, errors.New("terjadi kesalahan pada server")
		}
		votings = append(votings, v)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterasi rows voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	return votings, nil
}

// GetVotingResult mengambil rekap jumlah suara per kandidat
func (s *votingService) GetVotingResult() ([]types.VotingResult, error) {
	rows, err := s.db.Query(`
		SELECT 
			k.id_kandidat,
			k.nomor_urut,
			k.nama_ketua,
			k.nama_wakil,
			COUNT(v.id_voting) AS jumlah_suara
		FROM kandidat k
		LEFT JOIN voting v ON k.id_kandidat = v.id_kandidat
		GROUP BY k.id_kandidat, k.nomor_urut, k.nama_ketua, k.nama_wakil
		ORDER BY k.nomor_urut ASC`)
	if err != nil {
		log.Printf("Error query hasil voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	defer rows.Close()

	var results []types.VotingResult
	for rows.Next() {
		var r types.VotingResult
		if err := rows.Scan(&r.IDKandidat, &r.NomorUrut, &r.NamaKetua, &r.NamaWakil, &r.JumlahSuara); err != nil {
			log.Printf("Error scan row hasil voting: %v", err)
			return nil, errors.New("terjadi kesalahan pada server")
		}
		results = append(results, r)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterasi rows hasil voting: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	return results, nil
}

