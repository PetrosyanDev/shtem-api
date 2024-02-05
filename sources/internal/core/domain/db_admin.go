package domain

import (
	"time"
)

type AdminBase struct {
	ID        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminToken struct {
	Id        int64
	Token     string
	TokenHash []byte
	CreatedAt time.Time
	UpdatedAt time.Time
	Expiry    time.Time
	AdminId   int64
}

type Admin struct {
	AdminBase
	Username string
	Password string
	Token    AdminToken
}
