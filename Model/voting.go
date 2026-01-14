package models

import "time"

type Voting struct {
	IDVote    int64
	UserID    int64
	Kandidat  string
	Status    string
	CreatedAt time.Time
}
