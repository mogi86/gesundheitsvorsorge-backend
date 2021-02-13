package model

import (
	"time"
)

type User struct {
	ID                     uint64    `json:"id"`
	Password               string    `json:"password"`
	FirstName              string    `json:"first_name"`
	LastName               string    `json:"last_name"`
	Mail                   string    `json:"mail"`
	Sex                    string    `json:"sex"`
	Birthday               time.Time `json:"birthday"`
	Weight                 float64   `json:"weight"`
	Height                 float64   `json:"height"`
	Status                 bool      `json:"status"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
	TemporaryRegistrations []*TemporaryRegistration
}

type TemporaryRegistration struct {
	ID        uint64
	UserID    uint64
	Token     string
	ExpireAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
