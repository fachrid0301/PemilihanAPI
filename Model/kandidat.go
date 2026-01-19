package models

type Kandidat struct {
	IDKandidat int    `json:"id_kandidat"`
	NomorUrut  int    `json:"nomor_urut"`
	NamaKetua  string `json:"nama_ketua"`
	NamaWakil  string `json:"nama_wakil"`
	Visi       string `json:"visi"`
	Misi       string `json:"misi"`
	Foto       string `json:"foto"`
	CreatedAt  string `json:"created_at"`
}
