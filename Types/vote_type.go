package types

import "time"

type VoteData struct {
	IDVote    int64     `json:"id_vote"`
	UserID    int64     `json:"id_user"`
	Kandidat  string    `json:"kandidat"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type VoteRequest struct {
	UserID   int64  `json:"user_id"`
	Kandidat string `json:"kandidat"`
}
