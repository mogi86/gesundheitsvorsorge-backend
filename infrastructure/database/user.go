package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"

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
	tx, err := d.DB.Begin()
	if err != nil {
		logrus.Errorf("failed begin transaction. %+v\n", err)
	}

	ins := `
INSERT INTO users (
	password,
	first_name,
	last_name,
	mail,
	sex,
	birthday,
	height,
	weight,
	status,
	created_at,
	updated_at
) VALUES (
	?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`
	stmt, err := d.DB.Prepare(ins)
	if err != nil {
		logrus.Errorf("failed prepare statement. %+v\n", err)
	}

	result, err := stmt.Exec(
		m.Password,
		m.FirstName,
		m.LastName,
		m.Mail,
		m.Sex,
		m.Birthday,
		m.Height,
		m.Weight,
		m.Status,
		m.CreatedAt,
		m.UpdatedAt,
	)
	if err != nil {
		logrus.Errorf("failed execute insert operation. %+v\n", err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		logrus.Errorf("failed get last inserted id. %+v\n", err)
	}
	m.ID = uint64(lastID)

	insTemp := `
INSERT INTO temporary_registrations (
	user_id,
	token,
	expire_date,
	created_at,
	updated_at
) VALUES (
	?, ?, ?, ?, ?
)
`

	stmtTemp, err := d.DB.Prepare(insTemp)
	if err != nil {
		logrus.Errorf("failed prepare statement for temporary_registrations. %+v\n", err)
	}

	resultTemp, err := stmtTemp.Exec(
		m.ID,
		m.TemporaryRegistration.Token,
		m.TemporaryRegistration.ExpireAt,
		m.TemporaryRegistration.CreatedAt,
		m.TemporaryRegistration.UpdatedAt,
	)
	if err != nil {
		logrus.Errorf("failed execute insert operation. %+v\n", err)
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("failed commit transaction. %+v\n", err)
	}

	lastIDTemp, err := resultTemp.LastInsertId()
	if err != nil {
		logrus.Errorf("failed get last inserted id for temporary_registrations. %+v\n", err)
	}
	m.TemporaryRegistration.ID = uint64(lastIDTemp)

	return m, nil
}
