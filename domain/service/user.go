package service

import (
	"github.com/mogi86/gesundheitsvorsorge-backend/domain/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type UserServiceInterface interface {
	IsExists(mail string) (bool, error)
}

type UserService struct {
	repository repository.User
}

func NewUserService(r repository.User) *UserService {
	return &UserService{
		repository: r,
	}
}

func (u *UserService) IsExists(mail string) (bool, error) {
	user, err := u.repository.FindByMail(mail)
	if err != nil {
		logrus.Errorf("failed to check user exists: %+v", err)
		return false, xerrors.Errorf("failed to check user exists: %w", err)
	}
	if user == nil {
		return false, nil
	}

	return true, nil
}
