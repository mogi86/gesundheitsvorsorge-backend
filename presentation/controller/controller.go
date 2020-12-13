package controller

import (
	"fmt"
	"net/http"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	usecase usecase.Interface
}

func NewController(useCase usecase.Interface) *Controller {
	return &Controller{
		usecase: useCase,
	}
}

func (s *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := s.usecase.Sample()

	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, response)
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}
