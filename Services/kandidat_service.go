package services

import (
	"database/sql"
	"errors"
	"log"

	db "PemilihanAPI/DB"
	models "PemilihanAPI/Model"
	types "PemilihanAPI/Types"
)

// KandidatService adalah interface untuk kandidat service
type KandidatService interface {
	CreateKandidat(req types.CreateKandidatRequest) (*models.Kandidat, error)
	GetAllKandidat() ([]models.Kandidat, error)
	GetKandidatByID(id int) (*models.Kandidat, error)
	UpdateKandidat(id int, req types.UpdateKandidatRequest) (*models.Kandidat, error)
	DeleteKandidat(id int) error
}

type kandidatService struct {
	db *sql.DB
}

// NewKandidatService membuat instance baru dari KandidatService
func NewKandidatService() KandidatService {
	return &kandidatService{
		db: db.DB,
	}
}

// CreateKandidat membuat kandidat baru
func (s *kandidatService) CreateKandidat(req types.CreateKandidatRequest) (*models.Kandidat, error) {
	// Cek apakah nomor urut sudah ada
	var existCount int
	err := s.db.QueryRow("SELECT COUNT(1) FROM kandidat WHERE nomor_urut = ?", req.NomorUrut).Scan(&existCount)
	if err != nil {
		log.Printf("Error checking nomor urut: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	if existCount > 0 {
		return nil, errors.New("nomor urut sudah digunakan")
	}

	// Insert kandidat ke database
	// Handle foto: jika kosong, kirim NULL
	var fotoValue interface{}
	if req.Foto == "" {
		fotoValue = nil
	} else {
		fotoValue = req.Foto
	}
	
	result, err := s.db.Exec(
		"INSERT INTO kandidat (nomor_urut, nama_ketua, nama_wakil, visi, misi, foto) VALUES (?, ?, ?, ?, ?, ?)",
		req.NomorUrut, req.NamaKetua, req.NamaWakil, req.Visi, req.Misi, fotoValue,
	)

	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server: " + err.Error())
	}

	// Ambil ID yang baru saja diinsert
	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.New("terjadi kesalahan pada server")
	}

	// Ambil data kandidat yang baru dibuat
	kandidat, err := s.GetKandidatByID(int(id))
	if err != nil {
		return nil, err
	}

	return kandidat, nil
}

// GetAllKandidat mengambil semua kandidat
func (s *kandidatService) GetAllKandidat() ([]models.Kandidat, error) {
	rows, err := s.db.Query("SELECT id_kandidat, nomor_urut, nama_ketua, nama_wakil, visi, misi, foto, created_at FROM kandidat ORDER BY nomor_urut ASC")
	if err != nil {
		log.Printf("Error querying all kandidat: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}
	defer rows.Close()

	var kandidats []models.Kandidat
	for rows.Next() {
		var kandidat models.Kandidat
		var foto sql.NullString
		err := rows.Scan(&kandidat.IDKandidat, &kandidat.NomorUrut, &kandidat.NamaKetua, &kandidat.NamaWakil, &kandidat.Visi, &kandidat.Misi, &foto, &kandidat.CreatedAt)
		if err != nil {
			log.Printf("Error scanning kandidat row: %v", err)
			return nil, errors.New("terjadi kesalahan pada server")
		}
		if foto.Valid {
			kandidat.Foto = foto.String
		} else {
			kandidat.Foto = ""
		}
		kandidats = append(kandidats, kandidat)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating kandidat rows: %v", err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	return kandidats, nil
}

// GetKandidatByID mengambil kandidat berdasarkan ID
func (s *kandidatService) GetKandidatByID(id int) (*models.Kandidat, error) {
	var kandidat models.Kandidat
	var foto sql.NullString
	err := s.db.QueryRow(
		"SELECT id_kandidat, nomor_urut, nama_ketua, nama_wakil, visi, misi, foto, created_at FROM kandidat WHERE id_kandidat = ?",
		id,
	).Scan(&kandidat.IDKandidat, &kandidat.NomorUrut, &kandidat.NamaKetua, &kandidat.NamaWakil, &kandidat.Visi, &kandidat.Misi, &foto, &kandidat.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("kandidat tidak ditemukan")
		}
		log.Printf("Error getting kandidat by ID %d: %v", id, err)
		return nil, errors.New("terjadi kesalahan pada server")
	}

	if foto.Valid {
		kandidat.Foto = foto.String
	} else {
		kandidat.Foto = ""
	}

	return &kandidat, nil
}

// UpdateKandidat mengupdate kandidat
func (s *kandidatService) UpdateKandidat(id int, req types.UpdateKandidatRequest) (*models.Kandidat, error) {
	// Cek apakah kandidat ada
	_, err := s.GetKandidatByID(id)
	if err != nil {
		return nil, err
	}

	// Jika nomor urut diupdate, cek apakah nomor urut sudah digunakan oleh kandidat lain
	if req.NomorUrut > 0 {
		var existCount int
		err := s.db.QueryRow("SELECT COUNT(1) FROM kandidat WHERE nomor_urut = ? AND id_kandidat != ?", req.NomorUrut, id).Scan(&existCount)
		if err != nil {
			log.Printf("Error checking nomor urut for update: %v", err)
			return nil, errors.New("terjadi kesalahan pada server")
		}
		if existCount > 0 {
			return nil, errors.New("nomor urut sudah digunakan")
		}
	}

	// Build query update dinamis berdasarkan field yang diisi
	query := "UPDATE kandidat SET "
	var args []interface{}
	var updates []string

	if req.NomorUrut > 0 {
		updates = append(updates, "nomor_urut = ?")
		args = append(args, req.NomorUrut)
	}
	if req.NamaKetua != "" {
		updates = append(updates, "nama_ketua = ?")
		args = append(args, req.NamaKetua)
	}
	if req.NamaWakil != "" {
		updates = append(updates, "nama_wakil = ?")
		args = append(args, req.NamaWakil)
	}
	if req.Visi != "" {
		updates = append(updates, "visi = ?")
		args = append(args, req.Visi)
	}
	if req.Misi != "" {
		updates = append(updates, "misi = ?")
		args = append(args, req.Misi)
	}
	// Foto diupdate jika pointer tidak nil (bisa string kosong untuk menghapus foto)
	if req.Foto != nil {
		updates = append(updates, "foto = ?")
		// Jika string kosong, kirim NULL ke database
		if *req.Foto == "" {
			args = append(args, nil)
		} else {
			args = append(args, *req.Foto)
		}
	}

	if len(updates) == 0 {
		return nil, errors.New("tidak ada data yang diupdate")
	}

	query += updates[0]
	for i := 1; i < len(updates); i++ {
		query += ", " + updates[i]
	}
	query += " WHERE id_kandidat = ?"
	args = append(args, id)

	_, err = s.db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating kandidat ID %d: %v | Query: %s | Args: %v", id, err, query, args)
		return nil, errors.New("terjadi kesalahan pada server: " + err.Error())
	}

	// Ambil data kandidat yang sudah diupdate
	kandidat, err := s.GetKandidatByID(id)
	if err != nil {
		return nil, err
	}

	return kandidat, nil
}

// DeleteKandidat menghapus kandidat
func (s *kandidatService) DeleteKandidat(id int) error {
	// Cek apakah kandidat ada
	_, err := s.GetKandidatByID(id)
	if err != nil {
		return err
	}

	// Hapus kandidat
	_, err = s.db.Exec("DELETE FROM kandidat WHERE id_kandidat = ?", id)
	if err != nil {
		log.Printf("Error deleting kandidat ID %d: %v", id, err)
		return errors.New("terjadi kesalahan pada server")
	}

	return nil
}
