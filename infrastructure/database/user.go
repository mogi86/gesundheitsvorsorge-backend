package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
)

type DBClient struct {
	DB *sql.DB
}

type UserDto struct {
	ID                       uint64    `db:"id"`
	Password                 string    `db:"password"`
	FirstName                string    `db:"first_name"`
	LastName                 string    `db:"last_name"`
	Mail                     string    `db:"mail"`
	Sex                      string    `db:"sex"`
	Birthday                 time.Time `db:"birthday"`
	Weight                   float64   `db:"weight"`
	Height                   float64   `db:"height"`
	Status                   bool      `db:"status"`
	CreatedAt                time.Time `db:"created_at"`
	UpdatedAt                time.Time `db:"updated_at"`
	TemporaryRegistrationDto *TemporaryRegistrationDto
}

type TemporaryRegistrationDto struct {
	ID        uint64    `db:"temp_id"`
	UserID    uint64    `db:"temp_user_id"`
	Token     string    `db:"temp_token"`
	ExpireAt  time.Time `db:"temp_expire_at"`
	CreatedAt time.Time `db:"temp_created_at"`
	UpdatedAt time.Time `db:"temp_updated_at"`
}

// FindById return specified 1 user
func (d *DBClient) FindById(id uint64) (*model.User, error) {
	sel := `
SELECT
	u.id,
	u.password,
	u.first_name,
	u.last_name,
	u.mail,
	u.sex,
	u.birthday,
	u.height,
	u.weight,
	u.status,
	u.created_at,
	u.updated_at,
	tr.id AS temp_id,
	tr.user_id AS temp_user_id,
	tr.token AS temp_token,
	tr.expire_date AS temp_expire_at,
	tr.created_at AS temp_created_at,
	tr.updated_at AS temp_updated_at
FROM
	users u,
	temporary_registrations tr
WHERE
	u.id = tr.user_id
AND
	u.id = ?
`

	stmt, err := d.DB.Prepare(sel)
	if err != nil {
		logrus.Errorf("failed prepare statement. %+v\n", err)
	}

	var user UserDto
	var temp TemporaryRegistrationDto

	err = stmt.QueryRow(id).Scan(
		&user.ID,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Mail,
		&user.Sex,
		&user.Birthday,
		&user.Weight,
		&user.Height,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&temp.ID,
		&temp.UserID,
		&temp.Token,
		&temp.ExpireAt,
		&temp.CreatedAt,
		&temp.UpdatedAt,
	)

	return &model.User{
		ID:        user.ID,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Mail:      user.Mail,
		Sex:       user.Sex,
		Birthday:  user.Birthday,
		Weight:    user.Weight,
		Height:    user.Height,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		TemporaryRegistration: &model.TemporaryRegistration{
			ID:        temp.ID,
			UserID:    temp.UserID,
			Token:     temp.Token,
			ExpireAt:  temp.ExpireAt,
			CreatedAt: temp.CreatedAt,
			UpdatedAt: temp.UpdatedAt,
		},
	}, nil
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
