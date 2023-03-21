package entity

import "time"

type Message struct {
	Guid      string
	ClientId  int64
	Text      string
	CreatedAt time.Time
}
