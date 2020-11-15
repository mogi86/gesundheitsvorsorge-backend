package main

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/mogi86/gesundheitsvorsorge-backend/application/usecase"
	"github.com/mogi86/gesundheitsvorsorge-backend/presentation/controller"
)

func main() {
	provider := usecase.New(&usecase.UseCase{})
	cont := controller.NewController(provider.UseCaseInterface)
	http.Handle("/", cont)

	logrus.Infof("build server...")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		logrus.Errorf("build server failed. %v\n", err)
	}
}
