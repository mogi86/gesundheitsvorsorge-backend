package response

import "time"

type UserCreate struct {
	ID        uint64    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Mail      string    `json:"mail"`
	Sex       string    `json:"sex"`
	Birthday  time.Time `json:"birthday"`
	Weight    float64   `json:"weight"`
	Height    float64   `json:"height"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
