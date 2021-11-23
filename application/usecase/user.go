package usecase

import (
	customeerr "github.com/mogi86/gesundheitsvorsorge-backend/application/error"
	"github.com/mogi86/gesundheitsvorsorge-backend/application/helper"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/service"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"

	"github.com/mogi86/gesundheitsvorsorge-backend/domain/model"
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/repository"
)

type UserInterface interface {
	GetUserById(id uint64) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	Login(mail, password string) (*model.User, error)
}

type User struct {
	service    service.UserServiceInterface
	repository repository.User
}

func NewUserUseCase(s service.UserServiceInterface, r repository.User) *User {
	return &User{
		service:    s,
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
	isExists, err := u.service.IsExists(user.Mail)
	if err != nil {
		logrus.Errorf("failed check exists mail. %+v", err)
		return nil, xerrors.Errorf("failed check exists mail: %w", err)
	}
	if isExists {
		return nil, &customeerr.NotFoundErr{}
	}

	user, err = u.repository.Create(user)
	if err != nil {
		logrus.Errorf("failed persist user model. %+v", err)
		return nil, xerrors.Errorf("failed persist user model: %w", err)
	}

	return user, nil
}

func (u *User) Login(mail, password string) (*model.User, error) {
	user, err := u.repository.FindByMail(mail)
	if err != nil {
		logrus.Errorf("failed to get user model. %+v", err)
		return nil, xerrors.Errorf("failed to get user model: %w", err)
	}

	if helper.ConvertToHash(password) != user.Password {
		logrus.Errorf("invalid password. %+v", err)
		return nil, xerrors.Errorf("invalid password: %w", err)
	}

	return user, nil
}
