package response

import (
	"encoding/json"
	"time"
)

type User struct {
	ID                    uint64                 `json:"id"`
	FirstName             string                 `json:"first_name"`
	LastName              string                 `json:"last_name"`
	Mail                  string                 `json:"mail"`
	Sex                   string                 `json:"sex"`
	Birthday              CustomDate             `json:"birthday"`
	Weight                float64                `json:"weight"`
	Height                float64                `json:"height"`
	Status                bool                   `json:"status"`
	CreatedAt             CustomDateTime         `json:"created_at"`
	UpdatedAt             CustomDateTime         `json:"updated_at"`
	TemporaryRegistration *TemporaryRegistration `json:"temporary_registrations"`
}

type TemporaryRegistration struct {
	ID        uint64         `json:"id"`
	UserID    uint64         `json:"user_id"`
	Token     string         `json:"token"`
	ExpireAt  CustomDateTime `json:"expire_at"`
	CreatedAt CustomDateTime `json:"created_at"`
	UpdatedAt CustomDateTime `json:"updated_at"`
}

type CustomDate time.Time

func (c *CustomDate) MarshalJSON() ([]byte, error) {
	formatted := time.Time(*c).Format("2006-01-02")
	return json.Marshal(formatted)
}

type CustomDateTime time.Time

func (c *CustomDateTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(*c).Format("2006-01-02 15:04:05")
	return json.Marshal(formatted)
}
