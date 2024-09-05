package entity

import "time"

type BlockEntity struct {
	Number    int
	Hash      string
	Timestamp time.Time
}
