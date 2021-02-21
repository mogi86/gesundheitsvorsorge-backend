package usecase

import (
	"github.com/mogi86/gesundheitsvorsorge-backend/application/helper"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"

	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/repository"
)

type UserInterface interface {
	GetUserById(id uint64) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	Login(mail, password string) error
}

type User struct {
	repository repository.User
}

func NewUserUseCase(r repository.User) *User {
	return &User{
		repository: r,
	}
}

func (u *User) GetUserById(id uint64) (*model.User, error) {
	user, err := u.repository.FindById(id)
	if err != nil {
		logrus.Errorf("couldn't find user: %+v", err)
		return nil, xerrors.Errorf("couldn't find user: %w", err)
	}

	return user, nil
}

func (u *User) CreateUser(user *model.User) (*model.User, error) {
	user, err := u.repository.Create(user)
	if err != nil {
		logrus.Errorf("failed persist user model. %+v", err)
		return nil, xerrors.Errorf("failed persist user model: %w", err)
	}

	return user, nil
}

func (u *User) Login(mail, password string) error {
	user, err := u.repository.FindByMail(mail)
	if err != nil {
		logrus.Errorf("failed to get user model. %+v", err)
		return xerrors.Errorf("failed to get user model: %w", err)
	}

	if helper.ConvertToHash(password) != user.Password {
		logrus.Errorf("invalid password. %+v", err)
		return xerrors.Errorf("invalid password: %w", err)
	}

	return nil
}
