package controller

import (
	"fmt"
	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	provider *usecase.Provider
}

func NewController(useCase usecase.Interface) *Controller {
	return &Controller{
		provider: &usecase.Provider{
			UseCaseInterface: useCase,
		},
	}
}

func (s *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := s.provider.UseCaseInterface.Sample()

	w.Header().Set("Content-Type", "application/json")
	_, err := fmt.Fprintf(w, response)
	if err != nil {
		logrus.Errorf("return response failed. %v\n", err)
	}
}
