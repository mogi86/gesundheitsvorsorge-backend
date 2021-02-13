package usecase

import (
	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/repository"
)

type UserInterface interface {
	GetUserById(id uint64) *model.User
	CreateUser(user *model.User) *model.User
}

type User struct {
	repository repository.User
}

func NewUserUseCase(r repository.User) *User {
	return &User{
		repository: r,
	}
}

func (u *User) GetUserById(id uint64) *model.User {
	user, err := u.repository.FindById(id)
	if err != nil {
		logrus.Errorf("couldn't find user: %+v\n", err)
	}

	return user
}

func (u *User) CreateUser(user *model.User) *model.User {
	user, err := u.repository.Create(user)
	if err != nil {
		logrus.Errorf("failed persist user model: %+v\n", err)
	}

	return user
}
