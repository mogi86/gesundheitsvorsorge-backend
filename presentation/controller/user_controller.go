package controller

import (
	"fmt"
	"net/http"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	usecase usecase.UserInterface
}

func NewUserController(userUseCase usecase.UserInterface) *UserController {
	return &UserController{
		usecase: userUseCase,
	}
}

func (s *UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("request: %+v", r))

	user := s.usecase.GetUserById(1)

	w.Header().Set("Content-Type", "application/json")

	_, err := fmt.Fprintf(w, fmt.Sprintf("%+v", user))
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}
