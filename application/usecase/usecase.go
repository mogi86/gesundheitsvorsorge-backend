package usecase

import "fmt"

type Interface interface {
	Sample() string
}

type Provider struct {
	UseCaseInterface Interface
}

func New(useCaseInterface Interface) *Provider {
	return &Provider{
		UseCaseInterface: useCaseInterface,
	}
}

type UseCase struct {
	//TODO
}

func (u *UseCase) Sample() string {
	return fmt.Sprintf("this is usecase!")
}
