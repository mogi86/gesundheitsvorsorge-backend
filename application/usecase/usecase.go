package usecase

type Interface interface {
	Sample() string
}

type UseCase struct {
	//TODO
}

func (u *UseCase) Sample() string {
	return "this is usecase!"
}
