package repository

import "github.com/mogi86/gesundheitsvorsorge-backend/domain/model"

type User interface {
	FindById(id uint64) (*model.User, error)
	FindByMail(mail string) (*model.User, error)
	Create(*model.User) (*model.User, error)
}
