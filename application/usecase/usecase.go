package usecase

import "fmt"

type Interface interface {
	Sample() string
}

type UseCase struct {
	//TODO
}

func (u *UseCase) Sample() string {
	return fmt.Sprintf("this is usecase!")
}
