package entity

import "time"

type Client struct {
	Guid      string
	ClientId  int64
	FirstName string
	LastName  string
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
