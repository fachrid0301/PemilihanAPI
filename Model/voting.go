package models

// Voting merepresentasikan satu record pemilihan (satu user memilih satu kandidat)
type Voting struct {
	IDVoting    int    `json:"id_voting"`
	IDUser      int    `json:"id_user"`
	IDKandidat  int    `json:"id_kandidat"`
	WaktuVoting string `json:"waktu_voting"`
}

