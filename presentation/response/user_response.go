package response

import "time"

type UserCreate struct {
	ID                     uint64                 `json:"id"`
	FirstName              string                 `json:"first_name"`
	LastName               string                 `json:"last_name"`
	Mail                   string                 `json:"mail"`
	Sex                    string                 `json:"sex"`
	Birthday               time.Time              `json:"birthday"`
	Weight                 float64                `json:"weight"`
	Height                 float64                `json:"height"`
	Status                 bool                   `json:"status"`
	CreatedAt              time.Time              `json:"created_at"`
	UpdatedAt              time.Time              `json:"updated_at"`
	TemporaryRegistrations *TemporaryRegistration `json:"temporary_registrations"`
}

type TemporaryRegistration struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Token     string    `json:"token"`
	ExpireAt  time.Time `json:"expire_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
