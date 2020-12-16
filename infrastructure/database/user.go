package database

import (
	"database/sql"

	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
)

type DBClient struct {
	DB *sql.DB
}

// FindById return specified 1 user
func (d *DBClient) FindById(id uint64) (*model.User, error) {
	return &model.User{}, nil
}

// Create create new user
func (d *DBClient) Create(m *model.User) (*model.User, error) {
	return m, nil
}
