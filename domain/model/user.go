package model

import (
	"time"
)

type User struct {
	id                     uint64
	password               string
	firstName              string
	lastName               string
	mail                   string
	sex                    string
	birthday               time.Time
	weight                 float32
	height                 float32
	status                 bool
	createdAt              time.Time
	updatedAt              time.Time
	temporaryRegistrations *TemporaryRegistration
}

type TemporaryRegistration struct {
	id        uint64
	token     string
	expireAt  time.Time
	createdAt time.Time
	updatedAt time.Time
}
